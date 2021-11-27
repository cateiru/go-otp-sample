# Go OTP Sample

Goの`github.com/pquerna/otp`を使用したワンタイムパスワード生成実験

## 実験方法

- Google Authenticator
- Microsoft Authenticator
- 1password

以上、3つのワンタイムパスワード生成器を利用し、ハッシュアルゴリズム、パスコード長を変更しvalidateできるか検証する。

## 実験結果

|      テストアプリ       | デフォルト | アルゴリズム:SHA256 | パスコード長: 8 | secret: "hogehoge" |
| :---------------------: | :--------: | :-----------------: | :-------------: | :----------------: |
|  Google Authenticator   |     OK     |         No          |    No(8文字)    |         OK         |
| Microsoft Authenticator |     OK     |         OK          |   OK(6文字？)   |         OK         |
|        1password        |     OK     |         No          |    No(8文字)    |         OK         |

## まとめ

```go
ops := totp.GenerateOpts{
    Issuer:      "cateiru.com",
    AccountName: "Yuto Watanabe",
    Period:      30,
    SecretSize:  20,
    Secret:      []byte("hogehoge"),
    Digits:      otp.DigitsSix,
    Algorithm:   otp.AlgorithmSHA1,
    Rand:        rand.Reader,
}
```

- アルゴリズム: SHA1、パスコード長: 6のデフォルト設定値で利用するほうが良い。
- IssuerとAccountNameは生成APPに表示されるためちゃんとつけたほうが良い
- secretは環境変数などで管理して入力することでセキュアになる（はず）

## 参考文献

- https://zenn.dev/lapi/articles/2021-06-04-otp_tutorial
