package utils

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func verifyCert(certPEM string) error {
	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIEOzCCAyOgAwIBAgIUEk2zoXyeO22OCP5eYZ3fJsjCMmIwDQYJKoZIhvcNAQEL
BQAwgawxCzAJBgNVBAYTAkNOMREwDwYDVQQIDAhHdWkgWmhvdTERMA8GA1UEBwwI
R3VpIFlhbmcxNTAzBgNVBAoMLEd1aXpob3UgRmFyIEVhc3QgQ3JlZGl0IE1hbmFn
ZW1lbnQgQ28uLCBMdGQuMQ4wDAYDVQQLDAVJREZPUjEOMAwGA1UEAwwFaWRmb3Ix
IDAeBgkqhkiG9w0BCQEWEWlkZm9yQGZlLWNyZWQuY29tMB4XDTE5MTEwODA3MzYz
OVoXDTI5MTEwNTA3MzYzOVowgawxCzAJBgNVBAYTAkNOMREwDwYDVQQIDAhHdWkg
WmhvdTERMA8GA1UEBwwIR3VpIFlhbmcxNTAzBgNVBAoMLEd1aXpob3UgRmFyIEVh
c3QgQ3JlZGl0IE1hbmFnZW1lbnQgQ28uLCBMdGQuMQ4wDAYDVQQLDAVJREZPUjEO
MAwGA1UEAwwFaWRmb3IxIDAeBgkqhkiG9w0BCQEWEWlkZm9yQGZlLWNyZWQuY29t
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxF6FbpkUXWxGXhfW9n+A
Cp8Vg439MiFGvJhpkACz2qmOUFbswNF6P76KlCnGBs79FlbgwpNUR8T3qx2cSlnk
cYFqGg+PTdrqSre8LntbIbm6otnSth1s03ebNcKZLVkxSTeggB4pF8Ytv1zMT4+B
J0cbAsgjyqfRzkWeq7pQ96xgdGXu3IRV4n1FPnx6mFu7IEHpAd7GPs6pl73ipWek
3ZUCD9ScnkrTn5JGTxeMd8j0ODJ6Tg/QGaXEtfYnjpDIUDJ5LmfqQ3TB1BhUeCVZ
EKXsgt8xBEHhU6Fh00IPKrvUgCEO7pgcEedVhDmgixyxKgAiyhpgLFl/l/ZmX3Z+
vwIDAQABo1MwUTAdBgNVHQ4EFgQULKnkkJ/56q8QdZEJDaMc+gZt9JUwHwYDVR0j
BBgwFoAULKnkkJ/56q8QdZEJDaMc+gZt9JUwDwYDVR0TAQH/BAUwAwEB/zANBgkq
hkiG9w0BAQsFAAOCAQEAWcwGGEUYChjRXAMykLiE4q7LaQ4ZlQFuwVgosXEIKAsO
mUBIv2SD+cVrZu8Tp96VALRNfjemKyHLsaVmGyWALrQi3EnZFUgTvxvNR2bP8+bp
LdB3HZwnBrgFRlx2eBzXTPcA+mTA4aIoUscf1sFjdPkBp/w4guFd164AUU2eXMzN
CpfweSbUDiTKBnhUd+WCmYumKD+WQnHcSCEZ6+mk1zH0dQNMlKWRDbMlRjAk8+S1
WqNrQph6kf+ZipkVdJnYX8OpspxsypMkR1ACFnH+00rptDFCpUpiWVHfGZT4G0HG
R7SxiE0Pf25pDuRGQSYYtrGBP3u65m40sr+L5rzqgw==
-----END CERTIFICATE-----`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		return errors.New("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		return errors.New("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return errors.New("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		return errors.New("failed to verify certificate: " + err.Error())
	}

	return nil
}
