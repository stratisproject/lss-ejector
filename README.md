
# Masternode Staking Setup Guide

This guide will help you set up your validator and participate in Masternode Staking. By following these steps, you'll stake your STRAX tokens, receive liquid staking tokens (**mSTRAX**), and use them as collateral on the Masternode dApp—earning dual rewards as both a validator operator and a Masternode participant.

----------

## Overview of Components

Understanding the role of each component is key to a successful setup:

### 1. Ejector Service

The **Ejector Service** acts as an intermediary between your validator and the staking network. Its primary functions include:

-   **Withdrawal & Emergency Handling:** Managing withdrawals and addressing validator issues automatically.

### 2. Staking Deposit CLI Tool

The **staking-deposit-cli** tool is used for generating validator credentials and deposit data. Its key tasks are:

-   **Key Generation:** Creating a secure mnemonic and generating unique validator keys (keystores).
-   **Deposit Data Creation:** Producing JSON files that serve as proof of deposit and link your validator to the staking protocol.

### 3. Stratis Launcher

The **Stratis Launcher** simplifies the process of configuring and managing your validator host. You can set up your validator host using the latest release from the [stratis-node repository](https://github.com/stratisproject/stratis-node). It provides:

-   **User-Friendly Interface:** A simple platform for uploading validator keys and managing staking operations.
-   **Validator Management:** Automated configuration and integration of the required components.
-   **Monitoring & Updates:** Real-time insights into validator status, performance, and rewards.

These components work together to create a robust and secure environment for Masternode Staking.

----------

## Prerequisites

-   **SSH Access:** Ensure you have SSH access to your host.
-   **Basic Linux Knowledge:** Familiarity with terminal commands is recommended.
-   **STRAX Tokens:** You will need 1,000,000 STRAX for Mainnet staking (or 100,000 tSTRAX for testing via the [Stratis EVM Faucet](https://faucet.stratisevm.com/)).

----------

## Step 1: Setting Up Your Validator Host

Before using the Stratis Launcher, SSH into your host and follow these steps:

### A. Install System Dependencies

Update and upgrade your packages, then install the required dependencies:

```bash
apt update && apt upgrade -y
apt install -y git build-essential libssl-dev zlib1g-dev libbz2-dev \
  libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev \
  libncursesw5-dev xz-utils tk-dev libffi-dev liblzma-dev \
  software-properties-common supervisor
```

### B. Install PyEnv and Python 3.10

1.  **Clone PyEnv into your home directory:**
    
    ```bash
    git clone https://github.com/pyenv/pyenv.git ~/.pyenv
    
    ```
    
2.  **Configure environment variables.**  
    Add the following lines to your `~/.bashrc` or `~/.zshrc` file:
    
    ```bash
    export PYENV_ROOT="$HOME/.pyenv"
    export PATH="$PYENV_ROOT/bin:$PATH"
    eval "$(pyenv init --path)"
    eval "$(pyenv init -)"
    ```
    
3.  **Reload your shell:**
    
    ```bash
    source ~/.bashrc
    ```
    
4.  **Install Python 3.10 and set it as the global version:**
    
    ```bash
    pyenv install 3.10
    pyenv global 3.10
    ```
    

### C. Install Go 1.23.5

1.  **Download, extract, and configure Go:**
    
    ```bash
    wget https://go.dev/dl/go1.23.5.linux-amd64.tar.gz
    rm -rf /usr/local/go
    tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    source ~/.bashrc
    ```
    

### D. Clone Required Repositories

Clone the repositories that contain the tools needed for your validator:

```bash
git clone https://github.com/stratisproject/staking-deposit-cli
git clone https://github.com/stratisproject/lss-ejector
```

### E. Build and Install Components

#### Build staking-deposit-cli:

1.  Navigate to the `staking-deposit-cli` directory:
    
    ```bash
    cd staking-deposit-cli
    ```
    
2.  Build and install using:
    
    ```bash
    make build_linux
    ./deposit.sh install
    ```
    

#### Build the Ejector:

1.  Navigate to the ejector directory:
    
    ```bash
    cd ../lss-ejector/
    ```
    
2.  Build and install the ejector:
    
    ```bash
    make install
    ```
    

----------

## Step 2: Generate Validator Keys

1.  **Navigate to the staking-deposit-cli directory:**
    
    ```bash
    cd ../staking-deposit-cli
    ```
    
2.  **Run the key generation script with the execution address parameter:**  
    For the **beta network**, run:
    
    ```bash
    ./deposit.sh new-mnemonic --execution_address 0x00A6A303B2085857f04d700DF780dcEe72fc7048
    ```
    
    > **Note:** This execution address is specific to the beta network. For Mainnet, as the protocol is not deployed, there is no defined execution address.
    
3.  You will be prompted to provide the following inputs:
    
    -   **Number of Validators:**
        -   **Auroria:** The collateral requirement is 100,000 STRAX. Define **5 validators** (5 x 20,000 STRAX).
        -   **Mainnet:** The collateral requirement is 1,000,000 STRAX. Define **50 validators** (50 x 20,000 STRAX).
    -   **Withdrawal Recipient Address:**
        -   For the **beta network**, **set this value to:**  
            `0x00A6A303B2085857f04d700DF780dcEe72fc7048`
        -   **Important:** Do not change this value or use a personal address. For Mainnet, as the protocol is not deployed, there is no address to use. Failure to set this address correctly will result in activation failure.
    -   **Network/Chain Name:**
        -   Enter the correct network name (e.g., `auroria` for Auroria, `stratis` for Mainnet).
    -   **Password:**
        -   Choose a secure password to protect your validator keystore(s).
    
    **Note:** The field for "Amount a validator must deposit" is no longer needed and should be omitted.
    
4.  **Important:**
    
    -   Record the generated mnemonic phrase and password securely.
    -   The script will create a `validator_keys` directory and output `deposit_data.json`.
    - Keep these safe and copy to your local machine. Later we will upload the `validator_keys` via the Stratis Launcher and the `deposit_data.json` will be used to record your validators within the Masternode Staking protocol.

----------

## Step 3: Configure and Run Services

### A. Configure the Ejector Service Using Supervisor

1.  **Create an Ejector configuration file:**  
    Open or create the configuration file for Supervisor:
    
    ```bash
    vi /etc/supervisor/conf.d/ejector.conf
    ```
    
2.  **Paste the following configuration snippet (replace `<keystore_password>` with your password):**
    
    ```ini
    [program:lss-ejector]
    command=lss-ejector start \
      --execution_endpoint=http://localhost:8545 \
      --consensus_endpoint=http://localhost:3500 \
      --keys_dir=/root/staking-deposit-cli/validator_keys \
      --withdraw_address=0x00A6A303B2085857f04d700DF780dcEe72fc7048
    environment=KEYSTORE_PASSWORD=<keystore_password>
    redirect_stderr=true
    stdout_logfile=/var/log/supervisor-lss-ejector.log
    stdout_logfile_maxbytes=500MB
    stdout_logfile_backups=10
    autostart=true
    startretries=10
    stopwaitsecs=60
    ```
    **Important:**
    Be sure to set your password value here: `  environment=KEYSTORE_PASSWORD=<keystore_password>`
    
3.  **Reload Supervisor** to apply the new configuration:
    
    ```bash
    supervisorctl reread
    supervisorctl update
    ```
    

### B. Set Up Your Validator Host via Stratis Launcher

To set up your validator host, use the latest release from the [stratis-node repository](https://github.com/stratisproject/stratis-node). 

This Stratis Launcher streamlines the process of uploading your validator keys and managing staking operations. Download and install the appropriate release from the repository, connect to the host we've configured and then upload your keys via the Staking Tab.

----------

## Step 4: Deposit Collateral and Receive Liquid Tokens

1.  **Deposit Collateral:**
    
    -   Navigate to the [Masternode Staking Beta](https://beta.staking.masternode.stratisevm.com/).
    -   Deposit the required collateral (100,000 STRAX).
2.  **Receive Liquid Tokens:**
    
    -   Upon deposit confirmation, you will receive an equivalent amount of liquid Masternode Staking tokens (**mSTRAX**).
3.  **Upload Deposit Data:**
    
    -   Upload the `deposit_data` JSON file generated in Step 2 to link your validator with your staking deposit.

----------

## Step 5: Register on the Masternode dApp

1.  **Deposit mSTRAX Tokens:**
    
    -   Visit the Masternode dApp.
    -   Deposit your **mSTRAX** tokens as collateral to register on the Masternode Protocol.
2.  **Earn Dual Rewards:**
    
    -   Once registered, you will earn rewards both for operating your validator and for being a member of the Masternode contract.

----------

By following these detailed steps and understanding the role of each component, you'll be well-prepared to set up your validator host, configure essential services, and participate in Masternode Staking to maximize your yield.

**Happy Masternode Staking!**
