# Go-File-Encryption [WIP]

## How It Works
TODO


## Flags
Below is a list of the flags requried to execute the application. 

`-p`: the **password** you would like to encrypt/decrypt the file(s) with. 
    
`-path`: the **path** of the file(s) you would like to encrypt/decrypt.

`-e / -d`: indicate whether you would like to **encrypt** or **decrypt**. Exactly 1 of the flags must be provided.

## Example:

Encrypt:

    secureFiles -e -p "password" -path "./exampleFile.txt"
                            or
    secureFiles -e -p "password" -path "./exampleDirectory"

Decrypt:

    secureFiles -d -p "password" -path "./exampleFile.txt"
                            or
    secureFiles -d -p "password" -path "./exampleDirectory"
