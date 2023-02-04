
import "pe"

rule APT_NK_BabyShark_KimJoingRAT_Apr19_1 {
   meta:
      description = "Detects BabyShark KimJongRAT"
      author = "Florian Roth"
      reference = "https://unit42.paloaltonetworks.com/babyshark-malware-part-two-attacks-continue-using-kimjongrat-and-pcrat/"
      date = "2019-04-27"
      hash1 = "d50a0980da6297b8e4cec5db0a8773635cee74ac6f5c1ff18197dfba549f6712"
   strings:
      $x1 = "%s\\Microsoft\\ttmp.log" fullword wide

      $a1 = "logins.json" fullword ascii

      $s1 = "https://www.google.com/accounts/servicelogin" fullword ascii
      $s2 = "https://login.yahoo.com/config/login" fullword ascii
      $s3 = "SELECT id, hostname, httpRealm, formSubmitURL, usernameField, passwordField, encryptedUsername, encryptedPassword FROM moz_login" ascii
      $s4 = "\\mozsqlite3.dll" ascii
      $s5 = "SMTP Password" fullword ascii
      $s6 = "Yandex\\YandexBrowser\\User Data\\Default\\Login Data" fullword ascii
   condition:
      uint16(0) == 0x5a4d and filesize < 2000KB and (
         1 of ($x*) or
         ( $a1 and 3 of ($s*) )
      )
}
