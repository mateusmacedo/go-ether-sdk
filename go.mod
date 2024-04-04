module github.com/mateusmacedo/go-ether-sdk

go 1.22.0

replace github.com/mateusmacedo/go-ether-sdk/domain => ./pkg/domain

replace github.com/mateusmacedo/go-ether-sdk/application => ./pkg/application

require github.com/mateusmacedo/go-ether-sdk/domain v0.0.0-00010101000000-000000000000
