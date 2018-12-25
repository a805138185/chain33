// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockchain.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 区块头信息
// 	 version : 版本信息
// 	 parentHash :父哈希
// 	 txHash : 交易根哈希
// 	 stateHash :状态哈希
// 	 height : 区块高度
// 	 blockTime :区块产生时的时标
// 	 txCount : 区块上所有交易个数
// 	 difficulty :区块难度系数，
// 	 signature :交易签名
type Header struct {
	Version    int64      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte     `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte     `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash  []byte     `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height     int64      `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	BlockTime  int64      `protobuf:"varint,6,opt,name=blockTime" json:"blockTime,omitempty"`
	TxCount    int64      `protobuf:"varint,9,opt,name=txCount" json:"txCount,omitempty"`
	Hash       []byte     `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Difficulty uint32     `protobuf:"varint,11,opt,name=difficulty" json:"difficulty,omitempty"`
	Signature  *Signature `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Header) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Header) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Header) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Header) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *Header) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Header) GetDifficulty() uint32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Header) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

//  参考Header解释
type Block struct {
	Version    int64          `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte         `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte         `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	StateHash  []byte         `protobuf:"bytes,4,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Height     int64          `protobuf:"varint,5,opt,name=height" json:"height,omitempty"`
	BlockTime  int64          `protobuf:"varint,6,opt,name=blockTime" json:"blockTime,omitempty"`
	Difficulty uint32         `protobuf:"varint,11,opt,name=difficulty" json:"difficulty,omitempty"`
	Signature  *Signature     `protobuf:"bytes,8,opt,name=signature" json:"signature,omitempty"`
	Txs        []*Transaction `protobuf:"bytes,7,rep,name=txs" json:"txs,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Block) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Block) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Block) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *Block) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Block) GetDifficulty() uint32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Block) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Block) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Blocks struct {
	Items []*Block `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Blocks) Reset()                    { *m = Blocks{} }
func (m *Blocks) String() string            { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()               {}
func (*Blocks) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Blocks) GetItems() []*Block {
	if m != nil {
		return m.Items
	}
	return nil
}

type BlockSeqCB struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	URL    string `protobuf:"bytes,2,opt,name=URL" json:"URL,omitempty"`
	Encode string `protobuf:"bytes,3,opt,name=encode" json:"encode,omitempty"`
}

func (m *BlockSeqCB) Reset()                    { *m = BlockSeqCB{} }
func (m *BlockSeqCB) String() string            { return proto.CompactTextString(m) }
func (*BlockSeqCB) ProtoMessage()               {}
func (*BlockSeqCB) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *BlockSeqCB) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlockSeqCB) GetURL() string {
	if m != nil {
		return m.URL
	}
	return ""
}

func (m *BlockSeqCB) GetEncode() string {
	if m != nil {
		return m.Encode
	}
	return ""
}

type BlockSeqCBs struct {
	Items []*BlockSeqCB `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *BlockSeqCBs) Reset()                    { *m = BlockSeqCBs{} }
func (m *BlockSeqCBs) String() string            { return proto.CompactTextString(m) }
func (*BlockSeqCBs) ProtoMessage()               {}
func (*BlockSeqCBs) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *BlockSeqCBs) GetItems() []*BlockSeqCB {
	if m != nil {
		return m.Items
	}
	return nil
}

type BlockSeq struct {
	Num    int64          `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	Seq    *BlockSequence `protobuf:"bytes,2,opt,name=seq" json:"seq,omitempty"`
	Detail *BlockDetail   `protobuf:"bytes,3,opt,name=detail" json:"detail,omitempty"`
}

func (m *BlockSeq) Reset()                    { *m = BlockSeq{} }
func (m *BlockSeq) String() string            { return proto.CompactTextString(m) }
func (*BlockSeq) ProtoMessage()               {}
func (*BlockSeq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *BlockSeq) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *BlockSeq) GetSeq() *BlockSequence {
	if m != nil {
		return m.Seq
	}
	return nil
}

func (m *BlockSeq) GetDetail() *BlockDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

// 节点ID以及对应的Block
type BlockPid struct {
	Pid   string `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Block *Block `protobuf:"bytes,2,opt,name=block" json:"block,omitempty"`
}

func (m *BlockPid) Reset()                    { *m = BlockPid{} }
func (m *BlockPid) String() string            { return proto.CompactTextString(m) }
func (*BlockPid) ProtoMessage()               {}
func (*BlockPid) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *BlockPid) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *BlockPid) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

// resp
type BlockDetails struct {
	Items []*BlockDetail `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *BlockDetails) Reset()                    { *m = BlockDetails{} }
func (m *BlockDetails) String() string            { return proto.CompactTextString(m) }
func (*BlockDetails) ProtoMessage()               {}
func (*BlockDetails) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *BlockDetails) GetItems() []*BlockDetail {
	if m != nil {
		return m.Items
	}
	return nil
}

// resp
type Headers struct {
	Items []*Header `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Headers) Reset()                    { *m = Headers{} }
func (m *Headers) String() string            { return proto.CompactTextString(m) }
func (*Headers) ProtoMessage()               {}
func (*Headers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *Headers) GetItems() []*Header {
	if m != nil {
		return m.Items
	}
	return nil
}

type HeadersPid struct {
	Pid     string   `protobuf:"bytes,1,opt,name=pid" json:"pid,omitempty"`
	Headers *Headers `protobuf:"bytes,2,opt,name=headers" json:"headers,omitempty"`
}

func (m *HeadersPid) Reset()                    { *m = HeadersPid{} }
func (m *HeadersPid) String() string            { return proto.CompactTextString(m) }
func (*HeadersPid) ProtoMessage()               {}
func (*HeadersPid) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *HeadersPid) GetPid() string {
	if m != nil {
		return m.Pid
	}
	return ""
}

func (m *HeadersPid) GetHeaders() *Headers {
	if m != nil {
		return m.Headers
	}
	return nil
}

// 区块视图
// 	 head : 区块头信息
// 	 txCount :区块上交易个数
// 	 txHashes : 区块上交易的哈希列表
type BlockOverview struct {
	Head     *Header  `protobuf:"bytes,1,opt,name=head" json:"head,omitempty"`
	TxCount  int64    `protobuf:"varint,2,opt,name=txCount" json:"txCount,omitempty"`
	TxHashes [][]byte `protobuf:"bytes,3,rep,name=txHashes,proto3" json:"txHashes,omitempty"`
}

func (m *BlockOverview) Reset()                    { *m = BlockOverview{} }
func (m *BlockOverview) String() string            { return proto.CompactTextString(m) }
func (*BlockOverview) ProtoMessage()               {}
func (*BlockOverview) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *BlockOverview) GetHead() *Header {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *BlockOverview) GetTxCount() int64 {
	if m != nil {
		return m.TxCount
	}
	return 0
}

func (m *BlockOverview) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

// 区块详细信息
// 	 block : 区块信息
// 	 receipts :区块上所有交易的收据信息列表
type BlockDetail struct {
	Block          *Block         `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
	Receipts       []*ReceiptData `protobuf:"bytes,2,rep,name=receipts" json:"receipts,omitempty"`
	KV             []*KeyValue    `protobuf:"bytes,3,rep,name=KV" json:"KV,omitempty"`
	PrevStatusHash []byte         `protobuf:"bytes,4,opt,name=prevStatusHash,proto3" json:"prevStatusHash,omitempty"`
}

func (m *BlockDetail) Reset()                    { *m = BlockDetail{} }
func (m *BlockDetail) String() string            { return proto.CompactTextString(m) }
func (*BlockDetail) ProtoMessage()               {}
func (*BlockDetail) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

func (m *BlockDetail) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *BlockDetail) GetReceipts() []*ReceiptData {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func (m *BlockDetail) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

func (m *BlockDetail) GetPrevStatusHash() []byte {
	if m != nil {
		return m.PrevStatusHash
	}
	return nil
}

type Receipts struct {
	Receipts []*Receipt `protobuf:"bytes,1,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *Receipts) Reset()                    { *m = Receipts{} }
func (m *Receipts) String() string            { return proto.CompactTextString(m) }
func (*Receipts) ProtoMessage()               {}
func (*Receipts) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

func (m *Receipts) GetReceipts() []*Receipt {
	if m != nil {
		return m.Receipts
	}
	return nil
}

type PrivacyKV struct {
	PrivacyKVToken []*PrivacyKVToken `protobuf:"bytes,1,rep,name=privacyKVToken" json:"privacyKVToken,omitempty"`
}

func (m *PrivacyKV) Reset()                    { *m = PrivacyKV{} }
func (m *PrivacyKV) String() string            { return proto.CompactTextString(m) }
func (*PrivacyKV) ProtoMessage()               {}
func (*PrivacyKV) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

func (m *PrivacyKV) GetPrivacyKVToken() []*PrivacyKVToken {
	if m != nil {
		return m.PrivacyKVToken
	}
	return nil
}

type PrivacyKVToken struct {
	Token   string      `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	TxIndex int32       `protobuf:"varint,2,opt,name=txIndex" json:"txIndex,omitempty"`
	Txhash  []byte      `protobuf:"bytes,3,opt,name=txhash,proto3" json:"txhash,omitempty"`
	KV      []*KeyValue `protobuf:"bytes,4,rep,name=KV" json:"KV,omitempty"`
}

func (m *PrivacyKVToken) Reset()                    { *m = PrivacyKVToken{} }
func (m *PrivacyKVToken) String() string            { return proto.CompactTextString(m) }
func (*PrivacyKVToken) ProtoMessage()               {}
func (*PrivacyKVToken) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *PrivacyKVToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PrivacyKVToken) GetTxIndex() int32 {
	if m != nil {
		return m.TxIndex
	}
	return 0
}

func (m *PrivacyKVToken) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *PrivacyKVToken) GetKV() []*KeyValue {
	if m != nil {
		return m.KV
	}
	return nil
}

type ReceiptsAndPrivacyKV struct {
	Receipts  *Receipts  `protobuf:"bytes,1,opt,name=receipts" json:"receipts,omitempty"`
	PrivacyKV *PrivacyKV `protobuf:"bytes,2,opt,name=privacyKV" json:"privacyKV,omitempty"`
}

func (m *ReceiptsAndPrivacyKV) Reset()                    { *m = ReceiptsAndPrivacyKV{} }
func (m *ReceiptsAndPrivacyKV) String() string            { return proto.CompactTextString(m) }
func (*ReceiptsAndPrivacyKV) ProtoMessage()               {}
func (*ReceiptsAndPrivacyKV) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{15} }

func (m *ReceiptsAndPrivacyKV) GetReceipts() *Receipts {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func (m *ReceiptsAndPrivacyKV) GetPrivacyKV() *PrivacyKV {
	if m != nil {
		return m.PrivacyKV
	}
	return nil
}

type ReceiptCheckTxList struct {
	Errs []string `protobuf:"bytes,1,rep,name=errs" json:"errs,omitempty"`
}

func (m *ReceiptCheckTxList) Reset()                    { *m = ReceiptCheckTxList{} }
func (m *ReceiptCheckTxList) String() string            { return proto.CompactTextString(m) }
func (*ReceiptCheckTxList) ProtoMessage()               {}
func (*ReceiptCheckTxList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{16} }

func (m *ReceiptCheckTxList) GetErrs() []string {
	if m != nil {
		return m.Errs
	}
	return nil
}

// 区块链状态
// 	 currentHeight : 区块最新高度
// 	 mempoolSize :内存池大小
// 	 msgQueueSize : 消息队列大小
type ChainStatus struct {
	CurrentHeight int64 `protobuf:"varint,1,opt,name=currentHeight" json:"currentHeight,omitempty"`
	MempoolSize   int64 `protobuf:"varint,2,opt,name=mempoolSize" json:"mempoolSize,omitempty"`
	MsgQueueSize  int64 `protobuf:"varint,3,opt,name=msgQueueSize" json:"msgQueueSize,omitempty"`
}

func (m *ChainStatus) Reset()                    { *m = ChainStatus{} }
func (m *ChainStatus) String() string            { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()               {}
func (*ChainStatus) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{17} }

func (m *ChainStatus) GetCurrentHeight() int64 {
	if m != nil {
		return m.CurrentHeight
	}
	return 0
}

func (m *ChainStatus) GetMempoolSize() int64 {
	if m != nil {
		return m.MempoolSize
	}
	return 0
}

func (m *ChainStatus) GetMsgQueueSize() int64 {
	if m != nil {
		return m.MsgQueueSize
	}
	return 0
}

// 获取区块信息
// 	 start : 获取区块的开始高度
// 	 end :获取区块的结束高度
// 	 Isdetail : 是否需要获取区块的详细信息
// 	 pid : peer列表
type ReqBlocks struct {
	Start    int64    `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End      int64    `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	IsDetail bool     `protobuf:"varint,3,opt,name=isDetail" json:"isDetail,omitempty"`
	Pid      []string `protobuf:"bytes,4,rep,name=pid" json:"pid,omitempty"`
}

func (m *ReqBlocks) Reset()                    { *m = ReqBlocks{} }
func (m *ReqBlocks) String() string            { return proto.CompactTextString(m) }
func (*ReqBlocks) ProtoMessage()               {}
func (*ReqBlocks) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{18} }

func (m *ReqBlocks) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *ReqBlocks) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *ReqBlocks) GetIsDetail() bool {
	if m != nil {
		return m.IsDetail
	}
	return false
}

func (m *ReqBlocks) GetPid() []string {
	if m != nil {
		return m.Pid
	}
	return nil
}

type MempoolSize struct {
	Size int64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
}

func (m *MempoolSize) Reset()                    { *m = MempoolSize{} }
func (m *MempoolSize) String() string            { return proto.CompactTextString(m) }
func (*MempoolSize) ProtoMessage()               {}
func (*MempoolSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{19} }

func (m *MempoolSize) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ReplyBlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *ReplyBlockHeight) Reset()                    { *m = ReplyBlockHeight{} }
func (m *ReplyBlockHeight) String() string            { return proto.CompactTextString(m) }
func (*ReplyBlockHeight) ProtoMessage()               {}
func (*ReplyBlockHeight) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{20} }

func (m *ReplyBlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// 区块体信息
// 	 txs : 区块上所有交易列表
// 	 receipts :区块上所有交易的收据信息列表
type BlockBody struct {
	Txs      []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
	Receipts []*ReceiptData `protobuf:"bytes,2,rep,name=receipts" json:"receipts,omitempty"`
}

func (m *BlockBody) Reset()                    { *m = BlockBody{} }
func (m *BlockBody) String() string            { return proto.CompactTextString(m) }
func (*BlockBody) ProtoMessage()               {}
func (*BlockBody) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{21} }

func (m *BlockBody) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *BlockBody) GetReceipts() []*ReceiptData {
	if m != nil {
		return m.Receipts
	}
	return nil
}

//  区块追赶主链状态，用于判断本节点区块是否已经同步好
type IsCaughtUp struct {
	Iscaughtup bool `protobuf:"varint,1,opt,name=Iscaughtup" json:"Iscaughtup,omitempty"`
}

func (m *IsCaughtUp) Reset()                    { *m = IsCaughtUp{} }
func (m *IsCaughtUp) String() string            { return proto.CompactTextString(m) }
func (*IsCaughtUp) ProtoMessage()               {}
func (*IsCaughtUp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{22} }

func (m *IsCaughtUp) GetIscaughtup() bool {
	if m != nil {
		return m.Iscaughtup
	}
	return false
}

//  ntp时钟状态
type IsNtpClockSync struct {
	Isntpclocksync bool `protobuf:"varint,1,opt,name=isntpclocksync" json:"isntpclocksync,omitempty"`
}

func (m *IsNtpClockSync) Reset()                    { *m = IsNtpClockSync{} }
func (m *IsNtpClockSync) String() string            { return proto.CompactTextString(m) }
func (*IsNtpClockSync) ProtoMessage()               {}
func (*IsNtpClockSync) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{23} }

func (m *IsNtpClockSync) GetIsntpclocksync() bool {
	if m != nil {
		return m.Isntpclocksync
	}
	return false
}

type ChainExecutor struct {
	Driver    string `protobuf:"bytes,1,opt,name=driver" json:"driver,omitempty"`
	FuncName  string `protobuf:"bytes,2,opt,name=funcName" json:"funcName,omitempty"`
	StateHash []byte `protobuf:"bytes,3,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	Param     []byte `protobuf:"bytes,4,opt,name=param,proto3" json:"param,omitempty"`
	// 扩展字段，用于额外的用途
	Extra []byte `protobuf:"bytes,5,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (m *ChainExecutor) Reset()                    { *m = ChainExecutor{} }
func (m *ChainExecutor) String() string            { return proto.CompactTextString(m) }
func (*ChainExecutor) ProtoMessage()               {}
func (*ChainExecutor) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{24} }

func (m *ChainExecutor) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *ChainExecutor) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *ChainExecutor) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *ChainExecutor) GetParam() []byte {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *ChainExecutor) GetExtra() []byte {
	if m != nil {
		return m.Extra
	}
	return nil
}

//  通过block hash记录block的操作类型及add/del：1/2
type BlockSequence struct {
	Hash []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Type int64  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
}

func (m *BlockSequence) Reset()                    { *m = BlockSequence{} }
func (m *BlockSequence) String() string            { return proto.CompactTextString(m) }
func (*BlockSequence) ProtoMessage()               {}
func (*BlockSequence) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{25} }

func (m *BlockSequence) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockSequence) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

// resp
type BlockSequences struct {
	Items []*BlockSequence `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *BlockSequences) Reset()                    { *m = BlockSequences{} }
func (m *BlockSequences) String() string            { return proto.CompactTextString(m) }
func (*BlockSequences) ProtoMessage()               {}
func (*BlockSequences) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{26} }

func (m *BlockSequences) GetItems() []*BlockSequence {
	if m != nil {
		return m.Items
	}
	return nil
}

// 平行链区块详细信息
// 	 blockdetail : 区块详细信息
// 	 sequence :区块序列号
type ParaChainBlockDetail struct {
	Blockdetail *BlockDetail `protobuf:"bytes,1,opt,name=blockdetail" json:"blockdetail,omitempty"`
	Sequence    int64        `protobuf:"varint,2,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *ParaChainBlockDetail) Reset()                    { *m = ParaChainBlockDetail{} }
func (m *ParaChainBlockDetail) String() string            { return proto.CompactTextString(m) }
func (*ParaChainBlockDetail) ProtoMessage()               {}
func (*ParaChainBlockDetail) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{27} }

func (m *ParaChainBlockDetail) GetBlockdetail() *BlockDetail {
	if m != nil {
		return m.Blockdetail
	}
	return nil
}

func (m *ParaChainBlockDetail) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*Header)(nil), "types.Header")
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*Blocks)(nil), "types.Blocks")
	proto.RegisterType((*BlockSeqCB)(nil), "types.BlockSeqCB")
	proto.RegisterType((*BlockSeqCBs)(nil), "types.BlockSeqCBs")
	proto.RegisterType((*BlockSeq)(nil), "types.BlockSeq")
	proto.RegisterType((*BlockPid)(nil), "types.BlockPid")
	proto.RegisterType((*BlockDetails)(nil), "types.BlockDetails")
	proto.RegisterType((*Headers)(nil), "types.Headers")
	proto.RegisterType((*HeadersPid)(nil), "types.HeadersPid")
	proto.RegisterType((*BlockOverview)(nil), "types.BlockOverview")
	proto.RegisterType((*BlockDetail)(nil), "types.BlockDetail")
	proto.RegisterType((*Receipts)(nil), "types.Receipts")
	proto.RegisterType((*PrivacyKV)(nil), "types.PrivacyKV")
	proto.RegisterType((*PrivacyKVToken)(nil), "types.PrivacyKVToken")
	proto.RegisterType((*ReceiptsAndPrivacyKV)(nil), "types.ReceiptsAndPrivacyKV")
	proto.RegisterType((*ReceiptCheckTxList)(nil), "types.ReceiptCheckTxList")
	proto.RegisterType((*ChainStatus)(nil), "types.ChainStatus")
	proto.RegisterType((*ReqBlocks)(nil), "types.ReqBlocks")
	proto.RegisterType((*MempoolSize)(nil), "types.MempoolSize")
	proto.RegisterType((*ReplyBlockHeight)(nil), "types.ReplyBlockHeight")
	proto.RegisterType((*BlockBody)(nil), "types.BlockBody")
	proto.RegisterType((*IsCaughtUp)(nil), "types.IsCaughtUp")
	proto.RegisterType((*IsNtpClockSync)(nil), "types.IsNtpClockSync")
	proto.RegisterType((*ChainExecutor)(nil), "types.ChainExecutor")
	proto.RegisterType((*BlockSequence)(nil), "types.BlockSequence")
	proto.RegisterType((*BlockSequences)(nil), "types.BlockSequences")
	proto.RegisterType((*ParaChainBlockDetail)(nil), "types.ParaChainBlockDetail")
}

func init() { proto.RegisterFile("blockchain.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1097 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x97, 0xf3, 0xaf, 0xc9, 0x24, 0x0d, 0xbd, 0x55, 0x40, 0x51, 0x05, 0x5c, 0x6e, 0x39, 0x1d,
	0x51, 0x39, 0xa5, 0x52, 0x8b, 0x8e, 0x7b, 0x00, 0x09, 0x9a, 0x22, 0xb5, 0xd7, 0xe3, 0x28, 0xdb,
	0x5e, 0x1f, 0x78, 0xdb, 0xda, 0xdb, 0xda, 0x6a, 0xfc, 0xa7, 0xde, 0x75, 0x88, 0xf9, 0x0e, 0x7c,
	0x0a, 0xde, 0x10, 0x1f, 0x12, 0xed, 0xec, 0x3a, 0xb6, 0x43, 0x8b, 0x84, 0xc4, 0x0b, 0x6f, 0x3b,
	0xff, 0x67, 0x7e, 0xe3, 0x99, 0x31, 0xec, 0x5c, 0x2f, 0x62, 0xf7, 0xce, 0xf5, 0x79, 0x10, 0xcd,
	0x92, 0x34, 0x56, 0x31, 0x69, 0xab, 0x3c, 0x11, 0x72, 0xf7, 0x89, 0x4a, 0x79, 0x24, 0xb9, 0xab,
	0x82, 0xd8, 0x4a, 0x76, 0x07, 0x6e, 0x1c, 0x86, 0x05, 0x45, 0xff, 0x6c, 0x40, 0xe7, 0x44, 0x70,
	0x4f, 0xa4, 0x64, 0x0c, 0x5b, 0x4b, 0x91, 0xca, 0x20, 0x8e, 0xc6, 0xce, 0xc4, 0x99, 0x36, 0x59,
	0x41, 0x92, 0x4f, 0x01, 0x12, 0x9e, 0x8a, 0x48, 0x9d, 0x70, 0xe9, 0x8f, 0x1b, 0x13, 0x67, 0x3a,
	0x60, 0x15, 0x0e, 0xf9, 0x08, 0x3a, 0x6a, 0x85, 0xb2, 0x26, 0xca, 0x2c, 0x45, 0x3e, 0x86, 0x9e,
	0x54, 0x5c, 0x09, 0x14, 0xb5, 0x50, 0x54, 0x32, 0xb4, 0x95, 0x2f, 0x82, 0x5b, 0x5f, 0x8d, 0xdb,
	0x18, 0xce, 0x52, 0xda, 0x0a, 0xcb, 0xb9, 0x0c, 0x42, 0x31, 0xee, 0xa0, 0xa8, 0x64, 0xe8, 0x2c,
	0xd5, 0x6a, 0x1e, 0x67, 0x91, 0x1a, 0xf7, 0x4c, 0x96, 0x96, 0x24, 0x04, 0x5a, 0xbe, 0x0e, 0x04,
	0x18, 0x08, 0xdf, 0x3a, 0x73, 0x2f, 0xb8, 0xb9, 0x09, 0xdc, 0x6c, 0xa1, 0xf2, 0x71, 0x7f, 0xe2,
	0x4c, 0xb7, 0x59, 0x85, 0x43, 0x66, 0xd0, 0x93, 0xc1, 0x6d, 0xc4, 0x55, 0x96, 0x8a, 0x71, 0x77,
	0xe2, 0x4c, 0xfb, 0x07, 0x3b, 0x33, 0x84, 0x6e, 0x76, 0x51, 0xf0, 0x59, 0xa9, 0x42, 0x7f, 0x6f,
	0x40, 0xfb, 0x48, 0xe7, 0xf2, 0x3f, 0x41, 0xeb, 0x3f, 0xae, 0x9f, 0x3c, 0x87, 0xa6, 0x5a, 0xc9,
	0xf1, 0xd6, 0xa4, 0x39, 0xed, 0x1f, 0x10, 0xab, 0x79, 0x59, 0x7e, 0x63, 0x4c, 0x8b, 0xe9, 0x4b,
	0xe8, 0x20, 0x48, 0x92, 0x50, 0x68, 0x07, 0x4a, 0x84, 0x72, 0xec, 0xa0, 0xc5, 0xc0, 0x5a, 0xa0,
	0x94, 0x19, 0x11, 0x7d, 0x03, 0x80, 0xf4, 0x85, 0xb8, 0x9f, 0x1f, 0xe9, 0x2e, 0x46, 0x3c, 0x14,
	0x08, 0x6a, 0x8f, 0xe1, 0x9b, 0xec, 0x40, 0xf3, 0x3d, 0x7b, 0x8b, 0x50, 0xf6, 0x98, 0x7e, 0x6a,
	0x34, 0x44, 0xe4, 0xc6, 0x9e, 0x40, 0x0c, 0x7b, 0xcc, 0x52, 0xf4, 0x15, 0xf4, 0x4b, 0x5f, 0x92,
	0x7c, 0x5e, 0x0f, 0xff, 0xa4, 0x1a, 0x1e, 0x55, 0x8a, 0x1c, 0x12, 0xe8, 0x16, 0x4c, 0x1d, 0x2d,
	0xca, 0x42, 0xdb, 0x55, 0xfd, 0x24, 0x2f, 0xa0, 0x29, 0xc5, 0x3d, 0xc6, 0xef, 0x1f, 0x8c, 0x36,
	0x9c, 0x64, 0x22, 0x72, 0x05, 0xd3, 0x0a, 0x64, 0x0f, 0x3a, 0x9e, 0x50, 0x3c, 0x58, 0x60, 0x56,
	0x25, 0x40, 0xa8, 0x7a, 0x8c, 0x12, 0x66, 0x35, 0xe8, 0xb7, 0x36, 0xe2, 0x79, 0xe0, 0xe9, 0x88,
	0x49, 0xe0, 0xd9, 0x92, 0xf5, 0x53, 0xe3, 0x86, 0x4d, 0xb4, 0x31, 0x37, 0x70, 0x43, 0x11, 0x7d,
	0x0d, 0x83, 0x8a, 0x63, 0x49, 0xa6, 0xf5, 0x62, 0x1f, 0x0a, 0x6e, 0xab, 0x9d, 0xc1, 0x96, 0x99,
	0x79, 0x49, 0x3e, 0xab, 0x1b, 0x6d, 0x5b, 0x23, 0x23, 0x2e, 0xf4, 0x4f, 0x00, 0xac, 0xfe, 0xc3,
	0xd9, 0x4e, 0x61, 0xcb, 0x37, 0x72, 0x9b, 0xef, 0xb0, 0xe6, 0x46, 0xb2, 0x42, 0x4c, 0x7d, 0xd8,
	0xc6, 0x7c, 0x7e, 0x5c, 0x8a, 0x74, 0x19, 0x88, 0x5f, 0xc8, 0x33, 0x68, 0x69, 0x19, 0x7a, 0xfb,
	0x5b, 0x78, 0x14, 0x55, 0x27, 0xbe, 0x51, 0x9f, 0xf8, 0x5d, 0xe8, 0x9a, 0xd9, 0x11, 0x72, 0xdc,
	0x9c, 0x34, 0xa7, 0x03, 0xb6, 0xa6, 0xe9, 0x1f, 0x8e, 0xfd, 0x14, 0x4c, 0xe9, 0x25, 0xa2, 0xce,
	0xa3, 0x88, 0x92, 0x19, 0x74, 0x53, 0xe1, 0x8a, 0x20, 0x51, 0xba, 0x90, 0x2a, 0x88, 0xcc, 0xb0,
	0x8f, 0xb9, 0xe2, 0x6c, 0xad, 0x43, 0x9e, 0x42, 0xe3, 0xec, 0x0a, 0x23, 0xf7, 0x0f, 0x3e, 0xb0,
	0x9a, 0x67, 0x22, 0xbf, 0xe2, 0x8b, 0x4c, 0xb0, 0xc6, 0xd9, 0x15, 0x79, 0x01, 0xc3, 0x24, 0x15,
	0xcb, 0x0b, 0xc5, 0x55, 0x26, 0x2b, 0x73, 0xbd, 0xc1, 0xa5, 0xaf, 0xa0, 0xcb, 0x0a, 0xa7, 0x7b,
	0x95, 0x24, 0x4c, 0x53, 0x86, 0xf5, 0x24, 0xca, 0x04, 0xe8, 0x1b, 0xe8, 0x9d, 0xa7, 0xc1, 0x92,
	0xbb, 0xf9, 0xd9, 0x15, 0xf9, 0x46, 0x07, 0xb3, 0xc4, 0x65, 0x7c, 0x27, 0x22, 0x6b, 0xfe, 0xa1,
	0x35, 0x3f, 0xaf, 0x09, 0xd9, 0x86, 0x32, 0xcd, 0x61, 0x58, 0xd7, 0x20, 0x23, 0x68, 0x2b, 0xeb,
	0x47, 0xb7, 0xda, 0x10, 0xa6, 0x1d, 0xa7, 0x91, 0x27, 0x56, 0xd8, 0x8e, 0x36, 0x2b, 0x48, 0xb3,
	0xd8, 0xfc, 0xda, 0x62, 0xc3, 0x25, 0x6c, 0x60, 0x6a, 0x3d, 0x0a, 0x13, 0x95, 0x30, 0x2a, 0xca,
	0xff, 0x2e, 0xf2, 0xca, 0x8a, 0xbe, 0xa8, 0x41, 0xe1, 0x54, 0xcc, 0x0b, 0xf5, 0x4a, 0x33, 0x66,
	0xd0, 0x5b, 0x57, 0x64, 0x3f, 0xc3, 0x9d, 0xcd, 0xca, 0x59, 0xa9, 0x42, 0xa7, 0x40, 0xac, 0x97,
	0xb9, 0x2f, 0xdc, 0xbb, 0xcb, 0xd5, 0xdb, 0x40, 0xe2, 0x11, 0x11, 0x69, 0x6a, 0x90, 0xef, 0x31,
	0x7c, 0xd3, 0x1c, 0xfa, 0x73, 0x7d, 0x5a, 0x4d, 0xc3, 0xc8, 0x73, 0xd8, 0x76, 0xb3, 0x14, 0xd7,
	0xb9, 0x59, 0xc8, 0x66, 0x53, 0xd4, 0x99, 0x64, 0x02, 0xfd, 0x50, 0x84, 0x49, 0x1c, 0x2f, 0x2e,
	0x82, 0x5f, 0x85, 0xfd, 0x72, 0xab, 0x2c, 0x42, 0x61, 0x10, 0xca, 0xdb, 0x9f, 0x32, 0x91, 0x09,
	0x54, 0x69, 0xa2, 0x4a, 0x8d, 0x47, 0x39, 0xf4, 0x98, 0xb8, 0xb7, 0xcb, 0x74, 0x04, 0x6d, 0xa9,
	0x78, 0x5a, 0x04, 0x34, 0x84, 0x1e, 0x47, 0x11, 0x79, 0x36, 0x80, 0x7e, 0xea, 0xb1, 0x08, 0xe4,
	0x71, 0xb9, 0x88, 0xba, 0x6c, 0x4d, 0x17, 0xc3, 0xdb, 0xc2, 0xf2, 0xf4, 0x93, 0x3e, 0x83, 0xfe,
	0x0f, 0x95, 0xac, 0x08, 0xb4, 0xa4, 0xce, 0xc6, 0xc4, 0xc0, 0x37, 0xdd, 0x83, 0x1d, 0x26, 0x92,
	0x45, 0x8e, 0x79, 0xd8, 0xfa, 0xca, 0x7b, 0xe4, 0x54, 0xef, 0x91, 0xce, 0x18, 0xd5, 0x8e, 0x62,
	0x2f, 0x2f, 0xce, 0x85, 0xf3, 0x8f, 0xe7, 0xe2, 0xdf, 0x8e, 0x1d, 0x7d, 0x09, 0x70, 0x2a, 0xe7,
	0x3c, 0xbb, 0xf5, 0xd5, 0xfb, 0x44, 0x9f, 0xb8, 0x53, 0xe9, 0x22, 0x95, 0x25, 0x98, 0x4c, 0x97,
	0x55, 0x38, 0xf4, 0x35, 0x0c, 0x4f, 0xe5, 0x3b, 0x95, 0xcc, 0x71, 0x5f, 0xe7, 0x91, 0xab, 0xa7,
	0x32, 0x90, 0x91, 0x4a, 0x5c, 0x84, 0x35, 0x8f, 0x5c, 0x6b, 0xb5, 0xc1, 0xa5, 0xbf, 0x39, 0xb0,
	0x8d, 0x8d, 0xff, 0x7e, 0x25, 0xdc, 0x4c, 0xc5, 0xa9, 0x2e, 0xda, 0x4b, 0x83, 0xa5, 0x48, 0xed,
	0x48, 0x58, 0x4a, 0x23, 0x7e, 0x93, 0x45, 0xee, 0x3b, 0x7d, 0xb8, 0xcc, 0x95, 0x5a, 0xd3, 0xf5,
	0xb3, 0xde, 0xdc, 0x3c, 0xeb, 0x23, 0x68, 0x27, 0x3c, 0xe5, 0xa1, 0x5d, 0x0c, 0x86, 0xd0, 0x5c,
	0xb1, 0x52, 0x29, 0xc7, 0x5b, 0x3f, 0x60, 0x86, 0xa0, 0x5f, 0xd9, 0xe5, 0x59, 0x1c, 0x1d, 0xdd,
	0x2b, 0xf4, 0xea, 0x98, 0x3f, 0x1e, 0x74, 0x48, 0xa0, 0x75, 0x99, 0x27, 0xc5, 0x07, 0x87, 0x6f,
	0xfa, 0x35, 0x0c, 0x6b, 0x86, 0x7a, 0xc9, 0xd4, 0xd6, 0xfe, 0xc3, 0x37, 0xcd, 0x6e, 0x7f, 0x1f,
	0x46, 0xe7, 0x3c, 0xe5, 0x88, 0x44, 0x75, 0xa3, 0x7e, 0x09, 0x7d, 0x5c, 0x9b, 0xf6, 0xe4, 0x39,
	0x8f, 0x9e, 0xbc, 0xaa, 0x9a, 0x86, 0x4a, 0xda, 0x00, 0x36, 0xc7, 0x35, 0x7d, 0xf4, 0xf4, 0xe7,
	0x4f, 0x6e, 0x03, 0xe5, 0x67, 0xd7, 0x33, 0x37, 0x0e, 0xf7, 0x0f, 0x0f, 0xdd, 0x68, 0x1f, 0xff,
	0x69, 0x0f, 0x0f, 0xf7, 0xd1, 0xeb, 0x75, 0x07, 0x7f, 0x5a, 0x0f, 0xff, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x23, 0x1c, 0x7b, 0x8c, 0xf0, 0x0a, 0x00, 0x00,
}
