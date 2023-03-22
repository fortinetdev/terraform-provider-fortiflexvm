---
subcategory: ""
layout: "fortiflexvm"
page_title: "Generate an API token for FortiFlexVM"
description: |-
  Generate an API token for FortiFlexVM.
---

# Generate an API token for FortiFlexVM

FortiFlexVM Provider requires an API token to be authenticated. 

-> Before generating an API token for FortiFlexVM, please ensure you have a FortiCloud account and have FlexVM activated. 


* Step 1 (Optional): Create a new permission profile in IAM.

  Go to the [IAM Website](https://support.fortinet.com/iam/). Click `Permission Profiles` in the left navigation bar. On the new page, click `Add New` to create a new permission profile.

  On the `New Portal Permission Profile` page, fill in the `Permission Profile Name`, and keep `Status` as Active. 
  
  Click the `Add Portal` button. Select `FlexVM`. Then click the `Add` button.

  You will see `FlexVM` is listed under `PERMISSION PROFILE`. Click `Access` in `FlexVM` and set its `Access Type` as you want. Actions that involve changing or creating data (such as creating a new Configuration or updating a VM) will require ReadWrite permission or above.

  Click the `Submit` button in the upper right corner to submit.

* Step 2: Create an API User in IAM with permission to access FlexVM.

  In the [IAM Website](https://support.fortinet.com/iam/), click `Users` in the left navigation bar. On the new page, click `Add New > API User` to create an API User.

  In the `Select a Permission Profile`, select the user you created in Step 1. (If you skipped Step 1, you could select `SysAdmin`, in this case, you will create an admin user who has full access to Asset Management, IAM and FortiCare, which is not recommended).
  
  Click `Next > Confirm` to create an API user. The system will randomly assign a user name (API User ID).

* Step 3: Download your username and password.

  Go to the [IAM User Page](https://support.fortinet.com/iam/#/all-users). Click the user you created in Step 2.

  In the `API User Information` page, click the `Download Credentials` button in the bottom right to download your user name and password. 

  To prevent fraud, please set a password when you download your credential file and use this password to uncompress your credential file.

  **Downloading API User Credentials will reset the Users security credentials each time you perform this action.**

  If you download the credentials again, the previous password will become invalid.


Please refer [FlexVM Administration Guide](https://docs.fortinet.com/product/flex-vm/) for more information about FlexVM.