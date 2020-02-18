# Go-File-Encryption [WIP]

## How It Works

This is a basic file encryption script that utilizes the SHA256 hashing algorithm for computing the hash of the given password, which it then uses as the private key for 256-bit AES-GCM encryption.

**WARNING**: I do not recommend using this for sensitive data as it has not been thoroughly tested and was built as a learning exercise. 


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
