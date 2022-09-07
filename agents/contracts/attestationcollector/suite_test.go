package attestationcollector_test

import (
	"github.com/synapsecns/sanguine/core/testsuite"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
	"github.com/synapsecns/sanguine/agents/contracts/attestationcollector"
	"github.com/synapsecns/sanguine/agents/contracts/destination"
	"github.com/synapsecns/sanguine/agents/contracts/origin"
	"github.com/synapsecns/sanguine/agents/contracts/test/attestationharness"
	"github.com/synapsecns/sanguine/agents/testutil"
	"github.com/synapsecns/sanguine/ethergo/backends"
	"github.com/synapsecns/sanguine/ethergo/backends/preset"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
)

// AttestationCollectorSuite is the attestation collector test suite.
type AttestationCollectorSuite struct {
	*testsuite.TestSuite
	originContract              *origin.OriginRef
	destinationContract         *destination.DestinationRef
	destinationContractMetadata contracts.DeployedContract
	attestationHarness          *attestationharness.AttestationHarnessRef
	attestationContract         *attestationcollector.AttestationCollectorRef
	attestationContractMetadata contracts.DeployedContract
	testBackendOrigin           backends.SimulatedTestBackend
	testBackendDestination      backends.SimulatedTestBackend
	wallet                      wallet.Wallet
	signer                      signer.Signer
}

// NewAttestationCollectorSuite creates an end-to-end test suite.
func NewAttestationCollectorSuite(tb testing.TB) *AttestationCollectorSuite {
	tb.Helper()
	return &AttestationCollectorSuite{
		TestSuite: testsuite.NewTestSuite(tb),
	}
}

func (a *AttestationCollectorSuite) SetupTest() {
	a.TestSuite.SetupTest()

	a.testBackendOrigin = preset.GetRinkeby().Geth(a.GetTestContext(), a.T())
	a.testBackendDestination = preset.GetBSCTestnet().Geth(a.GetTestContext(), a.T())
	deployManager := testutil.NewDeployManager(a.T())

	_, a.originContract = deployManager.GetOrigin(a.GetTestContext(), a.testBackendOrigin)
	_, a.attestationHarness = deployManager.GetAttestationHarness(a.GetTestContext(), a.testBackendOrigin)
	a.attestationContractMetadata, a.attestationContract = deployManager.GetAttestationCollector(a.GetTestContext(), a.testBackendDestination)
	a.destinationContractMetadata, a.destinationContract = deployManager.GetDestination(a.GetTestContext(), a.testBackendDestination)

	var err error
	a.wallet, err = wallet.FromRandom()
	if err != nil {
		a.T().Fatal(err)
	}

	_, notaryManager := deployManager.GetNotaryManager(a.GetTestContext(), a.testBackendOrigin)
	owner, err := notaryManager.Owner(&bind.CallOpts{Context: a.GetTestContext()})
	if err != nil {
		a.T().Fatal(err)
	}

	a.signer = localsigner.NewSigner(a.wallet.PrivateKey())
	a.testBackendOrigin.FundAccount(a.GetTestContext(), a.signer.Address(), *big.NewInt(params.Ether))
	a.testBackendDestination.FundAccount(a.GetTestContext(), a.signer.Address(), *big.NewInt(params.Ether))

	transactOpts := a.testBackendOrigin.GetTxContext(a.GetTestContext(), &owner)

	tx, err := notaryManager.SetNotary(transactOpts.TransactOpts, a.signer.Address())
	if err != nil {
		a.T().Fatal(err)
	}

	a.testBackendOrigin.WaitForConfirmation(a.GetTestContext(), tx)
}

// TestAttestationCollectorSuite runs the integration test suite.
func TestAttestationCollectorSuite(t *testing.T) {
	suite.Run(t, NewAttestationCollectorSuite(t))
}