
$ gpg --list-keys
$ gpg --output theforce.gpg --encrypt --recipient 9186C4129FFD3D2500B35FA18E97CAEEEE861364 theforce.yaml

gpg --batch --yes --pinentry-mode loopback \
    --quick-generate-key \
    "Test Example <test@example.com>" \
    rsa2048 \
    sign,encrypt \
    0

gpg --armor --export test@example.com > pubkey.asc
gpg --armor --export-secret-keys test@example.com > privkey.asc

gpg --delete-secret-keys test@example.com
gpg --delete-keys test@example.com


gpg --show-keys gpg/pubkey.asc
gpg --import gpg/pubkey.asc
gpg --output vault/google.gpg --encrypt -r test@example.com seed/google.yaml
gpg --batch --yes --delete-keys test@example.com
