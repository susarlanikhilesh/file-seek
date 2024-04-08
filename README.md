# file-seek

File seek test to check if we can jump around the file using ReadOnly flag<br>
Tried the same in C also using Lseek

## Does not work with AppendOnly
Ref:  https://go-review.googlesource.com/c/go/+/14881 <br>
      https://go-review.googlesource.com/c/go/+/14807/2/src/os/os_test.go <br>  

## Os Info: 
> PRETTY_NAME="Ubuntu 22.04.4 LTS" <br>
> NAME="Ubuntu" <br>
> VERSION_ID="22.04" <br>
> VERSION="22.04.4 LTS (Jammy Jellyfish)" <br>
> VERSION_CODENAME=jammy <br>
> ID=ubuntu <br>
> ID_LIKE=debian <br>
> HOME_URL="https://www.ubuntu.com/" <br>
> SUPPORT_URL="https://help.ubuntu.com/" <br>
> BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/" <br>
> PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy" <br>
> UBUNTU_CODENAME=jammy <br>
