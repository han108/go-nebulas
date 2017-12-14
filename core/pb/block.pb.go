// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

/*
Package corepb is a generated protocol buffer package.

It is generated from these files:
	block.proto

It has these top-level messages:
	Account
	Data
	Transaction
	BlockHeader
	Block
	NetBlocks
	NetBlock
*/
package corepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Account struct {
	Balance    []byte `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Nonce      uint64 `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	VarsHash   []byte `protobuf:"bytes,3,opt,name=vars_hash,json=varsHash,proto3" json:"vars_hash,omitempty"`
	BirthPlace []byte `protobuf:"bytes,4,opt,name=birth_place,json=birthPlace,proto3" json:"birth_place,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{0} }

func (m *Account) GetBalance() []byte {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *Account) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Account) GetVarsHash() []byte {
	if m != nil {
		return m.VarsHash
	}
	return nil
}

func (m *Account) GetBirthPlace() []byte {
	if m != nil {
		return m.BirthPlace
	}
	return nil
}

type Data struct {
	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{1} }

func (m *Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Data) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Transaction struct {
	Hash      []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	From      []byte `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value     []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Nonce     uint64 `protobuf:"varint,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Timestamp int64  `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Data      *Data  `protobuf:"bytes,7,opt,name=data" json:"data,omitempty"`
	ChainId   uint32 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	GasPrice  []byte `protobuf:"bytes,9,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasLimit  []byte `protobuf:"bytes,10,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	Alg       uint32 `protobuf:"varint,11,opt,name=alg,proto3" json:"alg,omitempty"`
	Sign      []byte `protobuf:"bytes,12,opt,name=sign,proto3" json:"sign,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{2} }

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetFrom() []byte {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() []byte {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Transaction) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *Transaction) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *Transaction) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *Transaction) GetAlg() uint32 {
	if m != nil {
		return m.Alg
	}
	return 0
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

type BlockHeader struct {
	Hash       []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	ParentHash []byte `protobuf:"bytes,2,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Nonce      uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Coinbase   []byte `protobuf:"bytes,4,opt,name=coinbase,proto3" json:"coinbase,omitempty"`
	Timestamp  int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ChainId    uint32 `protobuf:"varint,6,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	StateRoot  []byte `protobuf:"bytes,7,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	TxsRoot    []byte `protobuf:"bytes,8,opt,name=txs_root,json=txsRoot,proto3" json:"txs_root,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{3} }

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetCoinbase() []byte {
	if m != nil {
		return m.Coinbase
	}
	return nil
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetChainId() uint32 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *BlockHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockHeader) GetTxsRoot() []byte {
	if m != nil {
		return m.TxsRoot
	}
	return nil
}

type Block struct {
	Header       *BlockHeader   `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Transactions []*Transaction `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Height       uint64         `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{4} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type NetBlocks struct {
	From   string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Batch  uint64   `protobuf:"varint,2,opt,name=batch,proto3" json:"batch,omitempty"`
	Blocks []*Block `protobuf:"bytes,3,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *NetBlocks) Reset()                    { *m = NetBlocks{} }
func (m *NetBlocks) String() string            { return proto.CompactTextString(m) }
func (*NetBlocks) ProtoMessage()               {}
func (*NetBlocks) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{5} }

func (m *NetBlocks) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *NetBlocks) GetBatch() uint64 {
	if m != nil {
		return m.Batch
	}
	return 0
}

func (m *NetBlocks) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type NetBlock struct {
	From  string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Batch uint64 `protobuf:"varint,2,opt,name=batch,proto3" json:"batch,omitempty"`
	Block *Block `protobuf:"bytes,3,opt,name=block" json:"block,omitempty"`
}

func (m *NetBlock) Reset()                    { *m = NetBlock{} }
func (m *NetBlock) String() string            { return proto.CompactTextString(m) }
func (*NetBlock) ProtoMessage()               {}
func (*NetBlock) Descriptor() ([]byte, []int) { return fileDescriptorBlock, []int{6} }

func (m *NetBlock) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *NetBlock) GetBatch() uint64 {
	if m != nil {
		return m.Batch
	}
	return 0
}

func (m *NetBlock) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "corepb.Account")
	proto.RegisterType((*Data)(nil), "corepb.Data")
	proto.RegisterType((*Transaction)(nil), "corepb.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "corepb.BlockHeader")
	proto.RegisterType((*Block)(nil), "corepb.Block")
	proto.RegisterType((*NetBlocks)(nil), "corepb.NetBlocks")
	proto.RegisterType((*NetBlock)(nil), "corepb.NetBlock")
}

func init() { proto.RegisterFile("block.proto", fileDescriptorBlock) }

var fileDescriptorBlock = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x8a, 0xd4, 0x4c,
	0x10, 0x25, 0x33, 0x93, 0x4c, 0x52, 0xc9, 0x7e, 0x7c, 0xb4, 0x8b, 0xb4, 0x7f, 0x6c, 0x88, 0x08,
	0x03, 0xc2, 0x5e, 0xac, 0x82, 0xd7, 0x8a, 0x17, 0x2b, 0x88, 0x2c, 0x8d, 0x37, 0x82, 0x30, 0x54,
	0x32, 0x71, 0x12, 0xcc, 0xa4, 0x43, 0xba, 0x76, 0xdd, 0x7d, 0x00, 0xdf, 0xc6, 0xf7, 0xf2, 0x35,
	0xa4, 0xab, 0x93, 0x99, 0x8c, 0x7a, 0xe3, 0x5d, 0x9d, 0x3a, 0xa9, 0xae, 0x3a, 0xa7, 0xab, 0x03,
	0x71, 0xde, 0xe8, 0xe2, 0xeb, 0x79, 0xd7, 0x6b, 0xd2, 0x22, 0x28, 0x74, 0x5f, 0x76, 0x79, 0xf6,
	0x0d, 0x96, 0xaf, 0x8b, 0x42, 0x5f, 0xb7, 0x24, 0x24, 0x2c, 0x73, 0x6c, 0xb0, 0x2d, 0x4a, 0xe9,
	0xa5, 0xde, 0x2a, 0x51, 0x23, 0x14, 0xa7, 0xe0, 0xb7, 0xda, 0xe6, 0x67, 0xa9, 0xb7, 0x5a, 0x28,
	0x07, 0xc4, 0x23, 0x88, 0x6e, 0xb0, 0x37, 0xeb, 0x0a, 0x4d, 0x25, 0xe7, 0x5c, 0x11, 0xda, 0xc4,
	0x25, 0x9a, 0x4a, 0x9c, 0x41, 0x9c, 0xd7, 0x3d, 0x55, 0xeb, 0xae, 0xc1, 0xa2, 0x94, 0x0b, 0xa6,
	0x81, 0x53, 0x57, 0x36, 0x93, 0xbd, 0x84, 0xc5, 0x5b, 0x24, 0x14, 0x02, 0x16, 0x74, 0xd7, 0xb9,
	0x96, 0x91, 0xe2, 0xd8, 0x4e, 0xd2, 0xe1, 0x5d, 0xa3, 0x71, 0xc3, 0x1d, 0x13, 0x35, 0xc2, 0xec,
	0xc7, 0x0c, 0xe2, 0x8f, 0x3d, 0xb6, 0x06, 0x0b, 0xaa, 0x75, 0x6b, 0xab, 0xb9, 0xbd, 0x1b, 0x98,
	0x63, 0x9b, 0xfb, 0xd2, 0xeb, 0xdd, 0x50, 0xca, 0xb1, 0xf8, 0x0f, 0x66, 0xa4, 0x87, 0x21, 0x67,
	0xa4, 0xad, 0xa2, 0x1b, 0x6c, 0xae, 0xc7, 0xc1, 0x1c, 0x38, 0xe8, 0xf4, 0xa7, 0x3a, 0x1f, 0x43,
	0x44, 0xf5, 0xae, 0x34, 0x84, 0xbb, 0x4e, 0x06, 0xa9, 0xb7, 0x9a, 0xab, 0x43, 0x42, 0xa4, 0xb0,
	0xd8, 0x20, 0xa1, 0x5c, 0xa6, 0xde, 0x2a, 0xbe, 0x48, 0xce, 0x9d, 0xaf, 0xe7, 0x56, 0x9b, 0x62,
	0x46, 0x3c, 0x80, 0xb0, 0xa8, 0xb0, 0x6e, 0xd7, 0xf5, 0x46, 0x86, 0xa9, 0xb7, 0x3a, 0x51, 0x4b,
	0xc6, 0xef, 0x36, 0xd6, 0xc2, 0x2d, 0x9a, 0x75, 0xd7, 0xd7, 0x45, 0x29, 0x23, 0x67, 0xe1, 0x16,
	0xcd, 0x95, 0xc5, 0x23, 0xd9, 0xd4, 0xbb, 0x9a, 0x24, 0xec, 0xc9, 0xf7, 0x16, 0x8b, 0xff, 0x61,
	0x8e, 0xcd, 0x56, 0xc6, 0x7c, 0x9e, 0x0d, 0xad, 0x6c, 0x53, 0x6f, 0x5b, 0x99, 0x38, 0xd9, 0x36,
	0xce, 0x7e, 0x7a, 0x10, 0xbf, 0xb1, 0xb7, 0x7e, 0x59, 0xe2, 0xa6, 0xec, 0xff, 0x6a, 0xd7, 0x19,
	0xc4, 0x1d, 0xf6, 0x65, 0x4b, 0xee, 0x22, 0x9d, 0x6b, 0xe0, 0x52, 0x7c, 0x95, 0x7b, 0x57, 0xe6,
	0x53, 0x57, 0x1e, 0x42, 0x58, 0xe8, 0xba, 0xcd, 0xd1, 0x8c, 0x26, 0xee, 0xf1, 0xb1, 0x63, 0xfe,
	0xef, 0x8e, 0x4d, 0xfd, 0x08, 0x8e, 0xfd, 0x78, 0x02, 0x60, 0x08, 0xa9, 0x5c, 0xf7, 0x5a, 0x13,
	0x5b, 0x9a, 0xa8, 0x88, 0x33, 0x4a, 0x6b, 0xb2, 0x95, 0x74, 0x6b, 0x1c, 0x19, 0xba, 0xc5, 0xa0,
	0x5b, 0x63, 0xa9, 0xec, 0xbb, 0x07, 0x3e, 0x2b, 0x15, 0xcf, 0x21, 0xa8, 0x58, 0x2d, 0xab, 0x8c,
	0x2f, 0xee, 0x8d, 0x57, 0x32, 0x31, 0x42, 0x0d, 0x9f, 0x88, 0x57, 0x90, 0xd0, 0x61, 0x9d, 0x8c,
	0x9c, 0xa5, 0xf3, 0x69, 0xc9, 0x64, 0xd5, 0xd4, 0xd1, 0x87, 0xe2, 0xbe, 0xed, 0x52, 0x6f, 0x2b,
	0x1a, 0x5c, 0x19, 0x50, 0xf6, 0x19, 0xa2, 0x0f, 0x25, 0x71, 0x2b, 0xb3, 0xdf, 0xc4, 0x61, 0xb7,
	0x79, 0x13, 0x4f, 0xc1, 0xcf, 0x91, 0x8a, 0x6a, 0x7c, 0x4b, 0x0c, 0xc4, 0x33, 0x08, 0xf8, 0x75,
	0x1a, 0x39, 0xe7, 0x09, 0x4e, 0x8e, 0x86, 0x56, 0x03, 0x99, 0x7d, 0x82, 0x70, 0x3c, 0xfd, 0x1f,
	0x0e, 0x7f, 0x0a, 0x3e, 0xd7, 0xf3, 0xa8, 0x7f, 0x9c, 0xed, 0xb8, 0x3c, 0xe0, 0xff, 0xc2, 0x8b,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0xd0, 0xb1, 0x12, 0x26, 0x04, 0x00, 0x00,
}
