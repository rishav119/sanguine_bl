package gasoracle

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../../../packages/contracts-core/flattened/GasOracle.sol --pkg gasoracle --sol-version 0.8.17 --filename gasoracle

// here we generate some interfaces we use in for our mocks. TODO this should be automated in abigen for all contracts + be condensed
//go:generate go run github.com/vburenin/ifacemaker -f gasoracle.abigen.go -s GasOracleCaller -i IGasOracleCaller -p gasoracle -o icaller_generated.go -c "autogenerated file"
//go:generate go run github.com/vburenin/ifacemaker -f gasoracle.abigen.go -s GasOracleTransactor -i IGasOracleTransactor -p gasoracle -o itransactor_generated.go -c "autogenerated file"
//go:generate go run github.com/vburenin/ifacemaker -f gasoracle.abigen.go -s GasOracleFilterer  -i IGasOracleFilterer  -p gasoracle  -o filterer_generated.go -c "autogenerated file"
//go:generate go run github.com/vektra/mockery/v2 --name IGasOracle --output ./mocks --case=underscore
// last line must be null