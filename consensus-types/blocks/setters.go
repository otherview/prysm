package blocks

import (
	"github.com/prysmaticlabs/prysm/v3/consensus-types/interfaces"
	"github.com/prysmaticlabs/prysm/v3/consensus-types/primitives"
	eth "github.com/prysmaticlabs/prysm/v3/proto/prysm/v1alpha1"
	"github.com/prysmaticlabs/prysm/v3/runtime/version"
)

// SetSignature sets the signature of the signed beacon block.
// This function is not thread safe, it is only used during block creation.
func (b *SignedBeaconBlock) SetSignature(sig []byte) {
	copy(b.signature[:], sig)
}

// SetSlot sets the respective slot of the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlock) SetSlot(slot primitives.Slot) {
	b.slot = slot
}

// SetProposerIndex sets the proposer index of the beacon block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlock) SetProposerIndex(proposerIndex primitives.ValidatorIndex) {
	b.proposerIndex = proposerIndex
}

// SetParentRoot sets the parent root of beacon block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlock) SetParentRoot(parentRoot []byte) {
	copy(b.parentRoot[:], parentRoot)
}

// SetStateRoot sets the state root of the underlying beacon block
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlock) SetStateRoot(root []byte) {
	copy(b.stateRoot[:], root)
}

// SetBlinded sets the blinded flag of the beacon block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlock) SetBlinded(blinded bool) {
	b.body.isBlinded = blinded
}

// SetRandaoReveal sets the randao reveal in the block body.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetRandaoReveal(r []byte) {
	copy(b.randaoReveal[:], r)
}

// SetGraffiti sets the graffiti in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetGraffiti(g []byte) {
	copy(b.graffiti[:], g)
}

// SetEth1Data sets the eth1 data in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetEth1Data(e *eth.Eth1Data) {
	b.eth1Data = e
}

// SetProposerSlashings sets the proposer slashings in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetProposerSlashings(p []*eth.ProposerSlashing) {
	b.proposerSlashings = p
}

// SetAttesterSlashings sets the attester slashings in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetAttesterSlashings(a []*eth.AttesterSlashing) {
	b.attesterSlashings = a
}

// SetAttestations sets the attestations in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetAttestations(a []*eth.Attestation) {
	b.attestations = a
}

// SetDeposits sets the deposits in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetDeposits(d []*eth.Deposit) {
	b.deposits = d
}

// SetVoluntaryExits sets the voluntary exits in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetVoluntaryExits(v []*eth.SignedVoluntaryExit) {
	b.voluntaryExits = v
}

// SetSyncAggregate sets the sync aggregate in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetSyncAggregate(s *eth.SyncAggregate) error {
	if b.version == version.Phase0 {
		return ErrNotSupported("SyncAggregate", b.version)
	}
	b.syncAggregate = s
	return nil
}

// SetExecution sets the execution payload of the block body.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetExecution(e interfaces.ExecutionData) error {
	if b.version == version.Phase0 || b.version == version.Altair {
		return ErrNotSupported("Execution", b.version)
	}
	if b.isBlinded {
		b.executionPayloadHeader = e
		return nil
	}
	b.executionPayload = e
	return nil
}

// SetBLSToExecutionChanges sets the BLS to execution changes in the block.
// This function is not thread safe, it is only used during block creation.
func (b *BeaconBlockBody) SetBLSToExecutionChanges(blsToExecutionChanges []*eth.SignedBLSToExecutionChange) error {
	if b.version < version.Capella {
		return ErrNotSupported("BLSToExecutionChanges", b.version)
	}
	b.blsToExecutionChanges = blsToExecutionChanges
	return nil
}
