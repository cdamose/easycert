// DO NOT EDIT THIS FILE. GENERATED BY "easycert-wrap help documentation".
// Edit the documentation in other files and re-run it to generate this one.

/*
EasyCert-wrap is a wrap over OpenSSL to create and handle certificates.

Usage:

        easycert-wrap command [arguments]

The commands are:

    init        initialize the directory
    ca          create certification authority
    req         create X509 certificate request
    sign        sign certificate request
    lang        generate files into a language to handle the certificate
    ls          list
    info        information
    cat         show the content
    chk         checking

Use "easycert-wrap help [command]" for more information about a command.


Initialize the directory

Usage:

        easycert-wrap init

"init" makes the directory structure in the HOME directory where
the certificates are handled.


Create certification authority

Usage:

        easycert-wrap ca [-rsa-size bits] [-years number]

"ca" creates a certification authority (CA) and makes the directories and files
to handle the certificates signed by this CA.


Create X509 certificate request

Usage:

        easycert-wrap req [-sign] [-rsa-size bits] [-years number] [-host name1,...] NAME

"req" creates a X509 certificate signing request (CSR) to be signed by a CA.


Sign certificate request

Usage:

        easycert-wrap sign [-years number] NAME

"sign" signs a certificate signing request (CSR) using the CA in the
certificates directory and generates a certificate.


Generate files into a language to handle the certificate

Usage:

        easycert-wrap lang [-ca file] [-server name] [-client] [-go]

"lang" generate files into a language to handle the certificate.
To look for the file, it uses the certificates directory when the "file" is just
a name or the path when the "file" is an absolute or relatative path.


List

Usage:

        easycert-wrap ls [-req] [-cert] [-key]

"ls" lists files in the certificates directory.
Whether it is not used some flag, it lists all files related to certificates.


Information

Usage:

        easycert-wrap info [-end-date] [-hash] [-issuer] [-name] FILE

"info" prints out information of a certificate.
To look for the file, it uses the certificates directory when the "file" is just
a name or the path when the "file" is an absolute or relatative path.

Whether a flag is not set, then it prints full information.


Show the content

Usage:

        easycert-wrap cat [-req | -cert | -key] FILE

"cat" shows the content of a certification-related file.
To look for the file, it uses the certificates directory when the "file" is just
a name or the path when the "file" is an absolute or relatative path.


Checking

Usage:

        easycert-wrap chk [-req | -cert | -key] FILE

"chk" checks whether a certification-related file is right.
To look for the file, it uses the certificates directory when the "file" is just
a name or the path when the "file" is an absolute or relatative path.


*/
package main
