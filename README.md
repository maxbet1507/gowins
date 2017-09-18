# gowins
golang for windows


## NonClientMetrics

[lxn/walk](https://github.com/lxn/walk)のデフォルトフォントがシステムフォントになっていないみたいなので、SystemParametersInfoから取得するための処理。

システム設定のDPIとか変更した際には再度`Get()`で取得しなおす必要があります。

> たしか変更された際にはWndProcとかで通知されるはず。
