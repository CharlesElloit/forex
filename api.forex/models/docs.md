## Password Policy Schema

| Field Name |Type| Description|
|------|-----|------|
| id   |  uuid   |   Unique identifier for the policy   |
| policy_name | string | Name of the password policy |
| description | string | Description of the policy |
| min_password_length | integer | Minimum length required for the password |
| max_password_length | integer | Maximum length allowed for the password |
| password_expiry_days | integer | Number of days before the password expires |
| grace_period_days | integer | Number of days before forced password change after expiration |
| enforce_password_history_limit | boolean | Enforces the password history limit to the define limit |
| password_history_limit | integer | Number of previous passwords remembered to prevent reuse |
| enable_complexity | boolean | Enforces password complexity |
| allow_repeated_characters | boolean | Enable or restricts the same characters from being repeated excessively |
| min_uppercase | integer | Minimum length allowed for uppercase letters |
| min_lowercase | integer | Minimum length allowed for the lowercase letters |
| min_numbers | integer | Minimum length allowed for the numbers in the password |
| min_special_characters | integer | Minimum length allowed for the special characters in the password |
| enforce_dictionary_check | boolean | Eabling a password check against a file.
| dictionary_source_file_path | string | A location to the file where the passwords will be check against |
| dictionary_source_api | string | The API which is hosting the password file |
| lockout_duration_minutes | integer | Duration (in minutes) for which the account remains locked |
| require_mfa_after_failed_attempts | integer | Enforces multi-factor authentication after failed login attempts |
| lockout_threshold | integer | Number of failed login attempts before account lockout |
| min_lockout_count | integer | A counter for how many time an account was locked |
| is_policy_active | boolean | Whether the policy is active or not |
| enable_permanent_lockout | boolean | Enforces a permanent lockout after serveral lockout |
| min_lockout_count | integer | Minimum number of account lockout before it's permanently lockout |
| created_date | date | date of the policy creation |
| created_time | time | Time of the policy creation |
| createdby_id | uuid | The identifier for the creator of the policy |
| lastmodified_date | date | The last date the policy was modified |
| lastmodified_time | time | The last time the policy was modified |
| lastmodifiedby_id | uuid | The identifier for the user who last modified the policy |
| source_ip | string | The source ip for the creator of the policy |
| source_browser | string | The browser or agent that created the policy |