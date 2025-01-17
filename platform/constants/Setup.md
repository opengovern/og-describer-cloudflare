# Setup GitHub Integration

Integrating opencomply with your GitHub organization using a Classic Personal Access Token (PAT) is a straightforward process. Here's a step-by-step guide to help you through it:

## Prerequisites

1. **opencomply Should be Installed and Running**: Ensure you have opencomply set up and ready to integrate with external services.
2. **GitHub Organization Admin Permissions**: Make sure you have administrative access to the GitHub organization you want to integrate with.
3. **Read Access to All Organization Repositories**: Ensure you have the necessary permissions to access the data in your GitHub organization.

<br>

## Create a Classic Personal Access Token (PAT)

1. **Access GitHub Settings**:
    - In the upper-right corner of any GitHub page, click your profile photo.
    - Select **Settings** from the dropdown menu.

2. **Go to Developer Settings**:
    - In the left sidebar, navigate to **Developer settings**.

3. **Generate a New Token**:
    - Under **Personal access tokens**, select **Tokens (classic)**.
    - Click **Generate new token**, and then **Generate new token (classic)** again.

4. **Set Token Expiration**:
    - Choose an expiration for the token by selecting **Expiration** and picking a default option or entering a custom date.

5. **Select Required Scopes**:
    - Ensure you select the following scopes to grant opencomply the necessary access:
        - `repo`
        - `read:org`
        - `read:packages`
        - `read:project`
        - `read:ssh_signing_key`
        - `read:audit_log`
        - `read:enterprise`
        - `read:discussion`
        - `read:user`
        - `user:email`
        - `notification`
        - `read:repo_hook`
        - `read:public_key`

6. **Generate the Token**:
    - After selecting the scopes, generate the token and copy it for use in opencomply integration.

## Configure Integration in opencomply

1. **Open opencomply Dashboard**:
    - Navigate to the opencomply dashboard.

2. **Go to Integrations**:
    - Select **Integrations** and then choose **GitHub** from the options.

3. **Select Classic PAT Integration**:
    - Opt for the **Classic PAT integration** method.

4. **Paste the PAT**:
    - Paste the copied PAT into the appropriate field.

5. **Select GitHub Organization**:
    - In the dropdown menu, choose the GitHub organization you want to connect.

6. **Finalize Integration**:
    - Click **Save** to confirm and enable the connection.

With these steps, opencomply will have read access to your GitHub repositories and related metadata, allowing it to provide governance and compliance oversight effectively.