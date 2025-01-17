# OpenComply Integration Setup with Cloudflare

To set up the OpenComply integration (formerly OpenGovernance) with Cloudflare, you need your **Account ID** and an **Account Token**. OpenComply requires a separate token for each Cloudflare account. If you're an enterprise with multiple accounts, ensure you create a token for each one.

## Prerequisites

- **Cloudflare Account** with necessary permissions.
- Access to [Cloudflare Dashboard](https://dash.cloudflare.com/).

## Steps to Create an Account API Token

### 1. Login to Cloudflare

- Visit [https://dash.cloudflare.com/](https://dash.cloudflare.com/).

### 2. Select Your Account

- If you have multiple accounts, click the **account selector** at the top of the dashboard.
- Choose the **account** for which you want to create the API token.

### 3. Copy Your Account ID

- On the **Account page**, locate your **Account ID**.
- **Copy** the Account ID and store it securely in a file. See screenshot below

![Copy Account ID](https://raw.githubusercontent.com/opengovern/hub/refs/heads/main/ui/src/pages/setup/setup-cloudflare.png)

> **Note:** Ensure you keep your Account ID confidential and store it in a secure location.

### 4. Navigate to Manage Account

- From the dashboard, click on **"Manage Account"** in the sidebar or dropdown menu.

### 5. Access Account API Tokens

- Within **Manage Account**, locate and click on **"Account API Tokens"**.

### 6. Create a New API Token

- Click the **"Create Token"** button to start the token creation process.

### 7. Provide Token Name

- In the **Name** field, enter: `opencomply`.

### 8. Select API Token Template

- Choose **"API token templates"**.
- Select **"Read all resources"** to grant necessary permissions.

### 9. Configure Zone Resources

- Under **Zone Resources**, select **"Include All zones from an account"**.

### 10. Set Token TTL

- Set the **TTL (Time to Live)** to **365 Days** to prevent integration issues.

### 11. Create the Token

- Click on **"Create Token"**.
- **Copy the Token Value immediately** as it won’t be shown again.

### 12. Securely Store the Token

- Save the **Token Value** in a secure location, such as a password manager.
- **Do not share** the token publicly to maintain account security.

## Provide Account ID and Token to OpenComply

Once you have both the **Account ID** and **Account Token**, follow these steps to integrate them with OpenComply:

1. **Log in to OpenComply** using your credentials.
2. Navigate to the **Integrations** section.
3. Select **Cloudflare Integration**.
4. **Enter** your **Account ID** and **Account Token** in the respective fields.
5. Click **Save** or **Connect** to finalize the integration.

## Summary

By following these steps, you will successfully create an **Account API Token** named **opencomply** on Cloudflare and obtain the **Token Value** and **Account ID** necessary for integrating with OpenComply. Ensure you store the token securely and set the TTL to 365 Days to maintain seamless integration.

---

If you encounter any issues, refer to [Cloudflare’s official documentation](https://developers.cloudflare.com/api/tokens/create/) or contact their support for assistance.

# OpenComply Support

For further assistance with the integration, please contact [OpenComply Support](mailto:support@opencomply.com).