# gowins
golang for windows


## NonClientMetrics

[lxn/walk](https://github.com/lxn/walk)のデフォルトフォントがシステムフォントになっていないみたいなので、SystemParametersInfoから取得するための処理。

`init()`時に取得していますが、システム設定のDPIとか変更した際には自動更新していないので`Update()`をコールする必要があります。

> たしか変更された際にはWndProcとかで通知されるはず。
