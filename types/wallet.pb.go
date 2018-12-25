// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 钱包模块存贮的tx交易详细信息
// 	 tx : tx交易信息
// 	 receipt :交易收据信息
// 	 height :交易所在的区块高度
// 	 index :交易所在区块中的索引
// 	 blocktime :交易所在区块的时标
// 	 amount :交易量
// 	 fromaddr :交易打出地址
// 	 txhash : 交易对应的哈希值
// 	 actionName  :交易对应的函数调用
//   payload: 保存额外的一些信息，主要是给插件使用
type WalletTxDetail struct {
	Tx         *Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
	Receipt    *ReceiptData `protobuf:"bytes,2,opt,name=receipt" json:"receipt,omitempty"`
	Height     int64        `protobuf:"varint,3,opt,name=height" json:"height,omitempty"`
	Index      int64        `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Blocktime  int64        `protobuf:"varint,5,opt,name=blocktime" json:"blocktime,omitempty"`
	Amount     int64        `protobuf:"varint,6,opt,name=amount" json:"amount,omitempty"`
	Fromaddr   string       `protobuf:"bytes,7,opt,name=fromaddr" json:"fromaddr,omitempty"`
	Txhash     []byte       `protobuf:"bytes,8,opt,name=txhash,proto3" json:"txhash,omitempty"`
	ActionName string       `protobuf:"bytes,9,opt,name=actionName" json:"actionName,omitempty"`
	Payload    []byte       `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *WalletTxDetail) Reset()                    { *m = WalletTxDetail{} }
func (m *WalletTxDetail) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetail) ProtoMessage()               {}
func (*WalletTxDetail) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *WalletTxDetail) GetTx() *Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *WalletTxDetail) GetReceipt() *ReceiptData {
	if m != nil {
		return m.Receipt
	}
	return nil
}

func (m *WalletTxDetail) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *WalletTxDetail) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *WalletTxDetail) GetBlocktime() int64 {
	if m != nil {
		return m.Blocktime
	}
	return 0
}

func (m *WalletTxDetail) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WalletTxDetail) GetFromaddr() string {
	if m != nil {
		return m.Fromaddr
	}
	return ""
}

func (m *WalletTxDetail) GetTxhash() []byte {
	if m != nil {
		return m.Txhash
	}
	return nil
}

func (m *WalletTxDetail) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *WalletTxDetail) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type WalletTxDetails struct {
	TxDetails []*WalletTxDetail `protobuf:"bytes,1,rep,name=txDetails" json:"txDetails,omitempty"`
}

func (m *WalletTxDetails) Reset()                    { *m = WalletTxDetails{} }
func (m *WalletTxDetails) String() string            { return proto.CompactTextString(m) }
func (*WalletTxDetails) ProtoMessage()               {}
func (*WalletTxDetails) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *WalletTxDetails) GetTxDetails() []*WalletTxDetail {
	if m != nil {
		return m.TxDetails
	}
	return nil
}

// 钱包模块存贮的账户信息
// 	 privkey : 账户地址对应的私钥
// 	 label :账户地址对应的标签
// 	 addr :账户地址
// 	 timeStamp :创建账户时的时标
type WalletAccountStore struct {
	Privkey   string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label     string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Addr      string `protobuf:"bytes,3,opt,name=addr" json:"addr,omitempty"`
	TimeStamp string `protobuf:"bytes,4,opt,name=timeStamp" json:"timeStamp,omitempty"`
}

func (m *WalletAccountStore) Reset()                    { *m = WalletAccountStore{} }
func (m *WalletAccountStore) String() string            { return proto.CompactTextString(m) }
func (*WalletAccountStore) ProtoMessage()               {}
func (*WalletAccountStore) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *WalletAccountStore) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *WalletAccountStore) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WalletAccountStore) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *WalletAccountStore) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

// 钱包模块通过一个随机值对钱包密码加密
// 	 pwHash : 对钱包密码和一个随机值组合进行哈希计算
// 	 randstr :对钱包密码加密的一个随机值
type WalletPwHash struct {
	PwHash  []byte `protobuf:"bytes,1,opt,name=pwHash,proto3" json:"pwHash,omitempty"`
	Randstr string `protobuf:"bytes,2,opt,name=randstr" json:"randstr,omitempty"`
}

func (m *WalletPwHash) Reset()                    { *m = WalletPwHash{} }
func (m *WalletPwHash) String() string            { return proto.CompactTextString(m) }
func (*WalletPwHash) ProtoMessage()               {}
func (*WalletPwHash) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *WalletPwHash) GetPwHash() []byte {
	if m != nil {
		return m.PwHash
	}
	return nil
}

func (m *WalletPwHash) GetRandstr() string {
	if m != nil {
		return m.Randstr
	}
	return ""
}

// 钱包当前的状态
// 	 isWalletLock : 钱包是否锁状态，true锁定，false解锁
// 	 isAutoMining :钱包是否开启挖矿功能，true开启挖矿，false关闭挖矿
// 	 isHasSeed : 钱包是否有种子，true已有，false没有
// 	 isTicketLock :钱包挖矿买票锁状态，true锁定，false解锁，只能用于挖矿转账
type WalletStatus struct {
	IsWalletLock bool `protobuf:"varint,1,opt,name=isWalletLock" json:"isWalletLock,omitempty"`
	IsAutoMining bool `protobuf:"varint,2,opt,name=isAutoMining" json:"isAutoMining,omitempty"`
	IsHasSeed    bool `protobuf:"varint,3,opt,name=isHasSeed" json:"isHasSeed,omitempty"`
	IsTicketLock bool `protobuf:"varint,4,opt,name=isTicketLock" json:"isTicketLock,omitempty"`
}

func (m *WalletStatus) Reset()                    { *m = WalletStatus{} }
func (m *WalletStatus) String() string            { return proto.CompactTextString(m) }
func (*WalletStatus) ProtoMessage()               {}
func (*WalletStatus) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *WalletStatus) GetIsWalletLock() bool {
	if m != nil {
		return m.IsWalletLock
	}
	return false
}

func (m *WalletStatus) GetIsAutoMining() bool {
	if m != nil {
		return m.IsAutoMining
	}
	return false
}

func (m *WalletStatus) GetIsHasSeed() bool {
	if m != nil {
		return m.IsHasSeed
	}
	return false
}

func (m *WalletStatus) GetIsTicketLock() bool {
	if m != nil {
		return m.IsTicketLock
	}
	return false
}

type WalletAccounts struct {
	Wallets []*WalletAccount `protobuf:"bytes,1,rep,name=wallets" json:"wallets,omitempty"`
}

func (m *WalletAccounts) Reset()                    { *m = WalletAccounts{} }
func (m *WalletAccounts) String() string            { return proto.CompactTextString(m) }
func (*WalletAccounts) ProtoMessage()               {}
func (*WalletAccounts) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *WalletAccounts) GetWallets() []*WalletAccount {
	if m != nil {
		return m.Wallets
	}
	return nil
}

type WalletAccount struct {
	Acc   *Account `protobuf:"bytes,1,opt,name=acc" json:"acc,omitempty"`
	Label string   `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *WalletAccount) Reset()                    { *m = WalletAccount{} }
func (m *WalletAccount) String() string            { return proto.CompactTextString(m) }
func (*WalletAccount) ProtoMessage()               {}
func (*WalletAccount) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *WalletAccount) GetAcc() *Account {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *WalletAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 钱包解锁
// 	 passwd : 钱包密码
// 	 timeout :钱包解锁时间，0，一直解锁，非0值，超时之后继续锁定
// 	 walletOrTicket :解锁整个钱包还是只解锁挖矿买票功能，1只解锁挖矿买票，0解锁整个钱包
type WalletUnLock struct {
	Passwd         string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
	Timeout        int64  `protobuf:"varint,2,opt,name=timeout" json:"timeout,omitempty"`
	WalletOrTicket bool   `protobuf:"varint,3,opt,name=walletOrTicket" json:"walletOrTicket,omitempty"`
}

func (m *WalletUnLock) Reset()                    { *m = WalletUnLock{} }
func (m *WalletUnLock) String() string            { return proto.CompactTextString(m) }
func (*WalletUnLock) ProtoMessage()               {}
func (*WalletUnLock) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *WalletUnLock) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *WalletUnLock) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *WalletUnLock) GetWalletOrTicket() bool {
	if m != nil {
		return m.WalletOrTicket
	}
	return false
}

type GenSeedLang struct {
	Lang int32 `protobuf:"varint,1,opt,name=lang" json:"lang,omitempty"`
}

func (m *GenSeedLang) Reset()                    { *m = GenSeedLang{} }
func (m *GenSeedLang) String() string            { return proto.CompactTextString(m) }
func (*GenSeedLang) ProtoMessage()               {}
func (*GenSeedLang) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *GenSeedLang) GetLang() int32 {
	if m != nil {
		return m.Lang
	}
	return 0
}

type GetSeedByPw struct {
	Passwd string `protobuf:"bytes,1,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *GetSeedByPw) Reset()                    { *m = GetSeedByPw{} }
func (m *GetSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*GetSeedByPw) ProtoMessage()               {}
func (*GetSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *GetSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

// 存储钱包的种子
// 	 seed : 钱包种子
// 	 passwd :钱包密码
type SaveSeedByPw struct {
	Seed   string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
	Passwd string `protobuf:"bytes,2,opt,name=passwd" json:"passwd,omitempty"`
}

func (m *SaveSeedByPw) Reset()                    { *m = SaveSeedByPw{} }
func (m *SaveSeedByPw) String() string            { return proto.CompactTextString(m) }
func (*SaveSeedByPw) ProtoMessage()               {}
func (*SaveSeedByPw) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *SaveSeedByPw) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *SaveSeedByPw) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ReplySeed struct {
	Seed string `protobuf:"bytes,1,opt,name=seed" json:"seed,omitempty"`
}

func (m *ReplySeed) Reset()                    { *m = ReplySeed{} }
func (m *ReplySeed) String() string            { return proto.CompactTextString(m) }
func (*ReplySeed) ProtoMessage()               {}
func (*ReplySeed) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *ReplySeed) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

type ReqWalletSetPasswd struct {
	OldPass string `protobuf:"bytes,1,opt,name=oldPass" json:"oldPass,omitempty"`
	NewPass string `protobuf:"bytes,2,opt,name=newPass" json:"newPass,omitempty"`
}

func (m *ReqWalletSetPasswd) Reset()                    { *m = ReqWalletSetPasswd{} }
func (m *ReqWalletSetPasswd) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetPasswd) ProtoMessage()               {}
func (*ReqWalletSetPasswd) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *ReqWalletSetPasswd) GetOldPass() string {
	if m != nil {
		return m.OldPass
	}
	return ""
}

func (m *ReqWalletSetPasswd) GetNewPass() string {
	if m != nil {
		return m.NewPass
	}
	return ""
}

type ReqNewAccount struct {
	Label string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
}

func (m *ReqNewAccount) Reset()                    { *m = ReqNewAccount{} }
func (m *ReqNewAccount) String() string            { return proto.CompactTextString(m) }
func (*ReqNewAccount) ProtoMessage()               {}
func (*ReqNewAccount) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *ReqNewAccount) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 获取钱包交易的详细信息
// 	 fromTx : []byte( Sprintf("%018d", height*100000 + index)，
// 				表示从高度 height 中的 index 开始获取交易列表；
// 			    第一次传参为空，获取最新的交易。)
// 	 count :获取交易列表的个数。
// 	 direction :查找方式；0，上一页；1，下一页。
type ReqWalletTransactionList struct {
	FromTx    []byte `protobuf:"bytes,1,opt,name=fromTx,proto3" json:"fromTx,omitempty"`
	Count     int32  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Direction int32  `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *ReqWalletTransactionList) Reset()                    { *m = ReqWalletTransactionList{} }
func (m *ReqWalletTransactionList) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletTransactionList) ProtoMessage()               {}
func (*ReqWalletTransactionList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{14} }

func (m *ReqWalletTransactionList) GetFromTx() []byte {
	if m != nil {
		return m.FromTx
	}
	return nil
}

func (m *ReqWalletTransactionList) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqWalletTransactionList) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type ReqWalletImportPrivkey struct {
	// bitcoin 的私钥格式
	Privkey string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Label   string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletImportPrivkey) Reset()                    { *m = ReqWalletImportPrivkey{} }
func (m *ReqWalletImportPrivkey) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletImportPrivkey) ProtoMessage()               {}
func (*ReqWalletImportPrivkey) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{15} }

func (m *ReqWalletImportPrivkey) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqWalletImportPrivkey) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// 发送交易
// 	 from : 打出地址
// 	 to :接受地址
// 	 amount : 转账额度
// 	 note :转账备注
type ReqWalletSendToAddress struct {
	From        string `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To          string `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Amount      int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note        []byte `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	IsToken     bool   `protobuf:"varint,5,opt,name=isToken" json:"isToken,omitempty"`
	TokenSymbol string `protobuf:"bytes,6,opt,name=tokenSymbol" json:"tokenSymbol,omitempty"`
}

func (m *ReqWalletSendToAddress) Reset()                    { *m = ReqWalletSendToAddress{} }
func (m *ReqWalletSendToAddress) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSendToAddress) ProtoMessage()               {}
func (*ReqWalletSendToAddress) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{16} }

func (m *ReqWalletSendToAddress) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqWalletSendToAddress) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqWalletSendToAddress) GetNote() []byte {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *ReqWalletSendToAddress) GetIsToken() bool {
	if m != nil {
		return m.IsToken
	}
	return false
}

func (m *ReqWalletSendToAddress) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

type ReqWalletSetFee struct {
	Amount int64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *ReqWalletSetFee) Reset()                    { *m = ReqWalletSetFee{} }
func (m *ReqWalletSetFee) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetFee) ProtoMessage()               {}
func (*ReqWalletSetFee) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{17} }

func (m *ReqWalletSetFee) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type ReqWalletSetLabel struct {
	Addr  string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
}

func (m *ReqWalletSetLabel) Reset()                    { *m = ReqWalletSetLabel{} }
func (m *ReqWalletSetLabel) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletSetLabel) ProtoMessage()               {}
func (*ReqWalletSetLabel) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{18} }

func (m *ReqWalletSetLabel) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqWalletSetLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type ReqWalletMergeBalance struct {
	To string `protobuf:"bytes,1,opt,name=to" json:"to,omitempty"`
}

func (m *ReqWalletMergeBalance) Reset()                    { *m = ReqWalletMergeBalance{} }
func (m *ReqWalletMergeBalance) String() string            { return proto.CompactTextString(m) }
func (*ReqWalletMergeBalance) ProtoMessage()               {}
func (*ReqWalletMergeBalance) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{19} }

func (m *ReqWalletMergeBalance) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type ReqTokenPreCreate struct {
	CreatorAddr  string `protobuf:"bytes,1,opt,name=creator_addr,json=creatorAddr" json:"creator_addr,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Symbol       string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
	Introduction string `protobuf:"bytes,4,opt,name=introduction" json:"introduction,omitempty"`
	OwnerAddr    string `protobuf:"bytes,5,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
	Total        int64  `protobuf:"varint,6,opt,name=total" json:"total,omitempty"`
	Price        int64  `protobuf:"varint,7,opt,name=price" json:"price,omitempty"`
}

func (m *ReqTokenPreCreate) Reset()                    { *m = ReqTokenPreCreate{} }
func (m *ReqTokenPreCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenPreCreate) ProtoMessage()               {}
func (*ReqTokenPreCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{20} }

func (m *ReqTokenPreCreate) GetCreatorAddr() string {
	if m != nil {
		return m.CreatorAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqTokenPreCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenPreCreate) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *ReqTokenPreCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

func (m *ReqTokenPreCreate) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReqTokenPreCreate) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type ReqTokenFinishCreate struct {
	FinisherAddr string `protobuf:"bytes,1,opt,name=finisher_addr,json=finisherAddr" json:"finisher_addr,omitempty"`
	Symbol       string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	OwnerAddr    string `protobuf:"bytes,3,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
}

func (m *ReqTokenFinishCreate) Reset()                    { *m = ReqTokenFinishCreate{} }
func (m *ReqTokenFinishCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenFinishCreate) ProtoMessage()               {}
func (*ReqTokenFinishCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{21} }

func (m *ReqTokenFinishCreate) GetFinisherAddr() string {
	if m != nil {
		return m.FinisherAddr
	}
	return ""
}

func (m *ReqTokenFinishCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenFinishCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

type ReqTokenRevokeCreate struct {
	RevokerAddr string `protobuf:"bytes,1,opt,name=revoker_addr,json=revokerAddr" json:"revoker_addr,omitempty"`
	Symbol      string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
	OwnerAddr   string `protobuf:"bytes,3,opt,name=owner_addr,json=ownerAddr" json:"owner_addr,omitempty"`
}

func (m *ReqTokenRevokeCreate) Reset()                    { *m = ReqTokenRevokeCreate{} }
func (m *ReqTokenRevokeCreate) String() string            { return proto.CompactTextString(m) }
func (*ReqTokenRevokeCreate) ProtoMessage()               {}
func (*ReqTokenRevokeCreate) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{22} }

func (m *ReqTokenRevokeCreate) GetRevokerAddr() string {
	if m != nil {
		return m.RevokerAddr
	}
	return ""
}

func (m *ReqTokenRevokeCreate) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *ReqTokenRevokeCreate) GetOwnerAddr() string {
	if m != nil {
		return m.OwnerAddr
	}
	return ""
}

type ReqModifyConfig struct {
	Key      string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Op       string `protobuf:"bytes,2,opt,name=op" json:"op,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Modifier string `protobuf:"bytes,4,opt,name=modifier" json:"modifier,omitempty"`
}

func (m *ReqModifyConfig) Reset()                    { *m = ReqModifyConfig{} }
func (m *ReqModifyConfig) String() string            { return proto.CompactTextString(m) }
func (*ReqModifyConfig) ProtoMessage()               {}
func (*ReqModifyConfig) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{23} }

func (m *ReqModifyConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ReqModifyConfig) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *ReqModifyConfig) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ReqModifyConfig) GetModifier() string {
	if m != nil {
		return m.Modifier
	}
	return ""
}

type ReqSignRawTx struct {
	Addr    string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Privkey string `protobuf:"bytes,2,opt,name=privkey" json:"privkey,omitempty"`
	TxHex   string `protobuf:"bytes,3,opt,name=txHex" json:"txHex,omitempty"`
	Expire  string `protobuf:"bytes,4,opt,name=expire" json:"expire,omitempty"`
	Index   int32  `protobuf:"varint,5,opt,name=index" json:"index,omitempty"`
	// 签名的模式类型
	// 0：普通交易
	// 1：隐私交易
	// int32  mode  = 6;
	Token     string `protobuf:"bytes,7,opt,name=token" json:"token,omitempty"`
	Fee       int64  `protobuf:"varint,8,opt,name=fee" json:"fee,omitempty"`
	NewExecer []byte `protobuf:"bytes,9,opt,name=newExecer,proto3" json:"newExecer,omitempty"`
	NewToAddr string `protobuf:"bytes,10,opt,name=newToAddr" json:"newToAddr,omitempty"`
}

func (m *ReqSignRawTx) Reset()                    { *m = ReqSignRawTx{} }
func (m *ReqSignRawTx) String() string            { return proto.CompactTextString(m) }
func (*ReqSignRawTx) ProtoMessage()               {}
func (*ReqSignRawTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{24} }

func (m *ReqSignRawTx) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *ReqSignRawTx) GetPrivkey() string {
	if m != nil {
		return m.Privkey
	}
	return ""
}

func (m *ReqSignRawTx) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

func (m *ReqSignRawTx) GetExpire() string {
	if m != nil {
		return m.Expire
	}
	return ""
}

func (m *ReqSignRawTx) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ReqSignRawTx) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ReqSignRawTx) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *ReqSignRawTx) GetNewExecer() []byte {
	if m != nil {
		return m.NewExecer
	}
	return nil
}

func (m *ReqSignRawTx) GetNewToAddr() string {
	if m != nil {
		return m.NewToAddr
	}
	return ""
}

type ReplySignRawTx struct {
	TxHex string `protobuf:"bytes,1,opt,name=txHex" json:"txHex,omitempty"`
}

func (m *ReplySignRawTx) Reset()                    { *m = ReplySignRawTx{} }
func (m *ReplySignRawTx) String() string            { return proto.CompactTextString(m) }
func (*ReplySignRawTx) ProtoMessage()               {}
func (*ReplySignRawTx) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{25} }

func (m *ReplySignRawTx) GetTxHex() string {
	if m != nil {
		return m.TxHex
	}
	return ""
}

type ReportErrEvent struct {
	Frommodule string `protobuf:"bytes,1,opt,name=frommodule" json:"frommodule,omitempty"`
	Tomodule   string `protobuf:"bytes,2,opt,name=tomodule" json:"tomodule,omitempty"`
	Error      string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *ReportErrEvent) Reset()                    { *m = ReportErrEvent{} }
func (m *ReportErrEvent) String() string            { return proto.CompactTextString(m) }
func (*ReportErrEvent) ProtoMessage()               {}
func (*ReportErrEvent) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{26} }

func (m *ReportErrEvent) GetFrommodule() string {
	if m != nil {
		return m.Frommodule
	}
	return ""
}

func (m *ReportErrEvent) GetTomodule() string {
	if m != nil {
		return m.Tomodule
	}
	return ""
}

func (m *ReportErrEvent) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Int32 struct {
	Data int32 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *Int32) Reset()                    { *m = Int32{} }
func (m *Int32) String() string            { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()               {}
func (*Int32) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{27} }

func (m *Int32) GetData() int32 {
	if m != nil {
		return m.Data
	}
	return 0
}

// 某些交易需要存在一些复杂的算法，所以需要请求服务器协助构建交易，返回对应交易哈希值，后续签名可以根据此哈希值进行处理
type ReqCreateTransaction struct {
	Tokenname string `protobuf:"bytes,1,opt,name=tokenname" json:"tokenname,omitempty"`
	// 构建交易类型
	// 0：普通的交易（暂不支持）
	// 1：隐私交易 公开->隐私
	// 2：隐私交易 隐私->隐私
	// 3：隐私交易 隐私->公开
	Type   int32  `protobuf:"varint,2,opt,name=type" json:"type,omitempty"`
	Amount int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
	Note   []byte `protobuf:"bytes,4,opt,name=note,proto3" json:"note,omitempty"`
	// 普通交易的发送方
	From string `protobuf:"bytes,5,opt,name=from" json:"from,omitempty"`
	// 普通交易的接收方
	To string `protobuf:"bytes,6,opt,name=to" json:"to,omitempty"`
	// 隐私交易，接收方的公钥对
	Pubkeypair string `protobuf:"bytes,10,opt,name=pubkeypair" json:"pubkeypair,omitempty"`
	Mixcount   int32  `protobuf:"varint,11,opt,name=mixcount" json:"mixcount,omitempty"`
	Expire     int64  `protobuf:"varint,12,opt,name=expire" json:"expire,omitempty"`
}

func (m *ReqCreateTransaction) Reset()                    { *m = ReqCreateTransaction{} }
func (m *ReqCreateTransaction) String() string            { return proto.CompactTextString(m) }
func (*ReqCreateTransaction) ProtoMessage()               {}
func (*ReqCreateTransaction) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{28} }

func (m *ReqCreateTransaction) GetTokenname() string {
	if m != nil {
		return m.Tokenname
	}
	return ""
}

func (m *ReqCreateTransaction) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqCreateTransaction) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ReqCreateTransaction) GetNote() []byte {
	if m != nil {
		return m.Note
	}
	return nil
}

func (m *ReqCreateTransaction) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ReqCreateTransaction) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ReqCreateTransaction) GetPubkeypair() string {
	if m != nil {
		return m.Pubkeypair
	}
	return ""
}

func (m *ReqCreateTransaction) GetMixcount() int32 {
	if m != nil {
		return m.Mixcount
	}
	return 0
}

func (m *ReqCreateTransaction) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type ReqAccountList struct {
	WithoutBalance bool `protobuf:"varint,1,opt,name=withoutBalance" json:"withoutBalance,omitempty"`
}

func (m *ReqAccountList) Reset()                    { *m = ReqAccountList{} }
func (m *ReqAccountList) String() string            { return proto.CompactTextString(m) }
func (*ReqAccountList) ProtoMessage()               {}
func (*ReqAccountList) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{29} }

func (m *ReqAccountList) GetWithoutBalance() bool {
	if m != nil {
		return m.WithoutBalance
	}
	return false
}

func init() {
	proto.RegisterType((*WalletTxDetail)(nil), "types.WalletTxDetail")
	proto.RegisterType((*WalletTxDetails)(nil), "types.WalletTxDetails")
	proto.RegisterType((*WalletAccountStore)(nil), "types.WalletAccountStore")
	proto.RegisterType((*WalletPwHash)(nil), "types.WalletPwHash")
	proto.RegisterType((*WalletStatus)(nil), "types.WalletStatus")
	proto.RegisterType((*WalletAccounts)(nil), "types.WalletAccounts")
	proto.RegisterType((*WalletAccount)(nil), "types.WalletAccount")
	proto.RegisterType((*WalletUnLock)(nil), "types.WalletUnLock")
	proto.RegisterType((*GenSeedLang)(nil), "types.GenSeedLang")
	proto.RegisterType((*GetSeedByPw)(nil), "types.GetSeedByPw")
	proto.RegisterType((*SaveSeedByPw)(nil), "types.SaveSeedByPw")
	proto.RegisterType((*ReplySeed)(nil), "types.ReplySeed")
	proto.RegisterType((*ReqWalletSetPasswd)(nil), "types.ReqWalletSetPasswd")
	proto.RegisterType((*ReqNewAccount)(nil), "types.ReqNewAccount")
	proto.RegisterType((*ReqWalletTransactionList)(nil), "types.ReqWalletTransactionList")
	proto.RegisterType((*ReqWalletImportPrivkey)(nil), "types.ReqWalletImportPrivkey")
	proto.RegisterType((*ReqWalletSendToAddress)(nil), "types.ReqWalletSendToAddress")
	proto.RegisterType((*ReqWalletSetFee)(nil), "types.ReqWalletSetFee")
	proto.RegisterType((*ReqWalletSetLabel)(nil), "types.ReqWalletSetLabel")
	proto.RegisterType((*ReqWalletMergeBalance)(nil), "types.ReqWalletMergeBalance")
	proto.RegisterType((*ReqTokenPreCreate)(nil), "types.ReqTokenPreCreate")
	proto.RegisterType((*ReqTokenFinishCreate)(nil), "types.ReqTokenFinishCreate")
	proto.RegisterType((*ReqTokenRevokeCreate)(nil), "types.ReqTokenRevokeCreate")
	proto.RegisterType((*ReqModifyConfig)(nil), "types.ReqModifyConfig")
	proto.RegisterType((*ReqSignRawTx)(nil), "types.ReqSignRawTx")
	proto.RegisterType((*ReplySignRawTx)(nil), "types.ReplySignRawTx")
	proto.RegisterType((*ReportErrEvent)(nil), "types.ReportErrEvent")
	proto.RegisterType((*Int32)(nil), "types.Int32")
	proto.RegisterType((*ReqCreateTransaction)(nil), "types.ReqCreateTransaction")
	proto.RegisterType((*ReqAccountList)(nil), "types.ReqAccountList")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 1305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5f, 0x6e, 0x1b, 0x37,
	0x13, 0xc7, 0x4a, 0x96, 0x6d, 0xd1, 0xb2, 0x93, 0x2c, 0x92, 0x40, 0xf0, 0xf7, 0x25, 0x71, 0x58,
	0x24, 0x75, 0x81, 0xc2, 0x01, 0xa2, 0x97, 0xa2, 0x40, 0x81, 0x38, 0x7f, 0x1d, 0xc0, 0x49, 0x0d,
	0x4a, 0x45, 0x81, 0xbe, 0x14, 0xd4, 0xee, 0x58, 0x22, 0xb4, 0x5a, 0xae, 0xb9, 0x94, 0x25, 0xdd,
	0xa4, 0x07, 0xe8, 0x11, 0x7a, 0x91, 0xde, 0xa1, 0x8f, 0x3d, 0x44, 0x31, 0x43, 0x52, 0xbb, 0x9b,
	0xb8, 0x0f, 0x41, 0xdf, 0xf8, 0x1b, 0x0e, 0x67, 0x38, 0xbf, 0x19, 0x0e, 0x87, 0xf5, 0x96, 0x32,
	0xcb, 0xc0, 0x9e, 0x14, 0x46, 0x5b, 0x1d, 0x77, 0xec, 0xba, 0x80, 0xf2, 0xf0, 0x8e, 0x35, 0x32,
	0x2f, 0x65, 0x62, 0x95, 0xce, 0xdd, 0xce, 0xe1, 0xed, 0x71, 0xa6, 0x93, 0x59, 0x32, 0x95, 0x2a,
	0x48, 0xf6, 0x65, 0x92, 0xe8, 0x45, 0xee, 0x8f, 0x1e, 0x1e, 0xc0, 0x0a, 0x92, 0x85, 0xd5, 0xc6,
	0x61, 0xfe, 0x47, 0x8b, 0x1d, 0xfc, 0x4c, 0xb6, 0x47, 0xab, 0xd7, 0x60, 0xa5, 0xca, 0x62, 0xce,
	0x5a, 0x76, 0xd5, 0x8f, 0x8e, 0xa2, 0xe3, 0xbd, 0xe7, 0xf1, 0x09, 0xb9, 0x3a, 0x19, 0x55, 0x9e,
	0x44, 0xcb, 0xae, 0xe2, 0x6f, 0xd9, 0x8e, 0x81, 0x04, 0x54, 0x61, 0xfb, 0xad, 0x86, 0xa2, 0x70,
	0xd2, 0xd7, 0xd2, 0x4a, 0x11, 0x54, 0xe2, 0xfb, 0x6c, 0x7b, 0x0a, 0x6a, 0x32, 0xb5, 0xfd, 0xf6,
	0x51, 0x74, 0xdc, 0x16, 0x1e, 0xc5, 0x77, 0x59, 0x47, 0xe5, 0x29, 0xac, 0xfa, 0x5b, 0x24, 0x76,
	0x20, 0xfe, 0x3f, 0xeb, 0x52, 0x14, 0x56, 0xcd, 0xa1, 0xdf, 0xa1, 0x9d, 0x4a, 0x80, 0xb6, 0xe4,
	0x1c, 0x03, 0xea, 0x6f, 0x3b, 0x5b, 0x0e, 0xc5, 0x87, 0x6c, 0xf7, 0xd2, 0xe8, 0xb9, 0x4c, 0x53,
	0xd3, 0xdf, 0x39, 0x8a, 0x8e, 0xbb, 0x62, 0x83, 0xf1, 0x8c, 0x5d, 0x4d, 0x65, 0x39, 0xed, 0xef,
	0x1e, 0x45, 0xc7, 0x3d, 0xe1, 0x51, 0xfc, 0x90, 0x31, 0x17, 0xd3, 0x47, 0x39, 0x87, 0x7e, 0x97,
	0x4e, 0xd5, 0x24, 0x71, 0x9f, 0xed, 0x14, 0x72, 0x9d, 0x69, 0x99, 0xf6, 0x19, 0x1d, 0x0c, 0x90,
	0xbf, 0x65, 0xb7, 0x9a, 0xac, 0x95, 0xf1, 0x80, 0x75, 0x6d, 0x00, 0xfd, 0xe8, 0xa8, 0x7d, 0xbc,
	0xf7, 0xfc, 0x9e, 0x27, 0xa5, 0xa9, 0x2a, 0x2a, 0x3d, 0x7e, 0xcd, 0x62, 0xb7, 0x79, 0xea, 0xb2,
	0x34, 0xb4, 0xda, 0x38, 0xbf, 0x46, 0x5d, 0xcf, 0x60, 0x4d, 0x69, 0xe8, 0x8a, 0x00, 0x91, 0xb1,
	0x4c, 0x8e, 0x21, 0x23, 0xd6, 0xbb, 0xc2, 0x81, 0x38, 0x66, 0x5b, 0x14, 0x77, 0x9b, 0x84, 0xb4,
	0x46, 0x16, 0x91, 0xaf, 0xa1, 0x95, 0xf3, 0x82, 0xf8, 0xed, 0x8a, 0x4a, 0xc0, 0x5f, 0xb0, 0x9e,
	0xf3, 0x7b, 0xb1, 0x3c, 0x43, 0x26, 0xee, 0xb3, 0xed, 0x82, 0x56, 0xe4, 0xb0, 0x27, 0x3c, 0xc2,
	0x9b, 0x18, 0x99, 0xa7, 0xa5, 0x35, 0xde, 0x63, 0x80, 0xfc, 0xb7, 0x28, 0x98, 0x18, 0x5a, 0x69,
	0x17, 0x65, 0xcc, 0x59, 0x4f, 0x95, 0x4e, 0x72, 0xae, 0x93, 0x19, 0x19, 0xda, 0x15, 0x0d, 0x99,
	0xd3, 0x39, 0x5d, 0x58, 0xfd, 0x41, 0xe5, 0x2a, 0x9f, 0x90, 0x4d, 0xd2, 0xa9, 0x64, 0x78, 0x71,
	0x55, 0x9e, 0xc9, 0x72, 0x08, 0x90, 0x52, 0x44, 0xbb, 0xa2, 0x12, 0x38, 0x0b, 0x23, 0x95, 0xcc,
	0xbc, 0x97, 0xad, 0x60, 0xa1, 0x92, 0xf1, 0x17, 0xa1, 0xa4, 0x3d, 0xa9, 0x65, 0x7c, 0xc2, 0x76,
	0xdc, 0x03, 0x0a, 0x99, 0xb9, 0xdb, 0xc8, 0x8c, 0xd7, 0x13, 0x41, 0x89, 0xbf, 0x63, 0xfb, 0x8d,
	0x9d, 0xf8, 0x88, 0xb5, 0x65, 0x92, 0xf8, 0x47, 0x71, 0xe0, 0x0f, 0x87, 0x63, 0xb8, 0x75, 0x73,
	0x66, 0xf8, 0x34, 0x90, 0xf4, 0x53, 0x4e, 0x04, 0x20, 0xcf, 0xb2, 0x2c, 0x97, 0xa9, 0x4f, 0xac,
	0x47, 0xc8, 0x33, 0x26, 0x47, 0x2f, 0xdc, 0x7b, 0x6a, 0x8b, 0x00, 0xe3, 0xa7, 0xec, 0xc0, 0xdd,
	0xea, 0x47, 0xe3, 0x42, 0xf4, 0x9c, 0x7c, 0x22, 0xe5, 0x8f, 0xd9, 0xde, 0x3b, 0xc8, 0x91, 0xa3,
	0x73, 0x99, 0x4f, 0xb0, 0x24, 0x32, 0x99, 0x4f, 0xc8, 0x4d, 0x47, 0xd0, 0x9a, 0x3f, 0x41, 0x15,
	0x8b, 0x2a, 0x2f, 0xd7, 0x17, 0xcb, 0x7f, 0xbb, 0x0b, 0xff, 0x9e, 0xf5, 0x86, 0xf2, 0x1a, 0x36,
	0x7a, 0x31, 0xdb, 0x2a, 0x31, 0x17, 0x4e, 0x8b, 0xd6, 0xb5, 0xb3, 0xad, 0xc6, 0xd9, 0x47, 0xac,
	0x2b, 0xa0, 0xc8, 0xd6, 0x94, 0xab, 0x1b, 0x0e, 0xf2, 0x33, 0x16, 0x0b, 0xb8, 0xf2, 0x85, 0x03,
	0xf6, 0x62, 0x13, 0xbe, 0xce, 0x52, 0x04, 0xa1, 0xe0, 0x3d, 0xc4, 0x9d, 0x1c, 0x96, 0xb4, 0xe3,
	0x0b, 0xd0, 0x43, 0xfe, 0x84, 0xed, 0x0b, 0xb8, 0xfa, 0x08, 0xcb, 0x90, 0xa3, 0x4d, 0x06, 0xa2,
	0x7a, 0x06, 0x2e, 0x59, 0x7f, 0xe3, 0xb0, 0xd6, 0xc5, 0xce, 0x55, 0x49, 0x7d, 0x09, 0x7b, 0xc4,
	0x68, 0x15, 0xaa, 0xde, 0x21, 0xb4, 0x44, 0x26, 0xc9, 0x65, 0x47, 0x38, 0x80, 0x85, 0x99, 0x2a,
	0x03, 0x74, 0x9c, 0x92, 0xd0, 0x11, 0x95, 0x80, 0x9f, 0xb1, 0xfb, 0x1b, 0x3f, 0xef, 0xe7, 0x85,
	0x36, 0xf6, 0xc2, 0xbf, 0xd9, 0x2f, 0x7c, 0xcd, 0xfc, 0xf7, 0xa8, 0x66, 0x6a, 0x08, 0x79, 0x3a,
	0xd2, 0xa7, 0x69, 0x6a, 0xa0, 0x2c, 0x91, 0x51, 0xbc, 0x62, 0x60, 0x14, 0xd7, 0xf1, 0x01, 0x6b,
	0x59, 0xed, 0x2d, 0xb4, 0xac, 0xae, 0x35, 0xc8, 0x76, 0xa3, 0x41, 0xc6, 0x6c, 0x2b, 0xd7, 0x16,
	0xe8, 0xc5, 0xf4, 0x04, 0xad, 0xf1, 0x6a, 0xaa, 0x1c, 0xe9, 0x19, 0xe4, 0xd4, 0x68, 0x77, 0x45,
	0x80, 0xf1, 0x11, 0xdb, 0xb3, 0xb8, 0x18, 0xae, 0xe7, 0x63, 0x9d, 0x51, 0xaf, 0xed, 0x8a, 0xba,
	0x88, 0x7f, 0xc3, 0x6e, 0xd5, 0x33, 0xf9, 0x16, 0xea, 0xbd, 0x39, 0xaa, 0xbb, 0xe6, 0x3f, 0xb0,
	0x3b, 0x75, 0xd5, 0xf3, 0x46, 0xd3, 0x8a, 0x6a, 0x4d, 0xeb, 0x66, 0x42, 0xbe, 0x66, 0xf7, 0x36,
	0xc7, 0x3f, 0x80, 0x99, 0xc0, 0x4b, 0x99, 0xc9, 0x3c, 0x01, 0x1f, 0x7a, 0x14, 0x42, 0xe7, 0x7f,
	0x46, 0xe4, 0x88, 0x22, 0xb8, 0x30, 0xf0, 0xca, 0x80, 0xb4, 0x10, 0x3f, 0x66, 0xbd, 0x04, 0x57,
	0xda, 0xfc, 0x5a, 0x73, 0xb8, 0xe7, 0x65, 0x48, 0x2d, 0x71, 0x83, 0x5f, 0x80, 0x73, 0x4b, 0x6b,
	0x0c, 0xa6, 0x74, 0xc1, 0xbb, 0xb6, 0xea, 0x11, 0x75, 0xa0, 0xdc, 0x1a, 0x9d, 0x2e, 0x5c, 0x25,
	0xb8, 0xde, 0xda, 0x90, 0xc5, 0x0f, 0x18, 0xd3, 0xcb, 0x1c, 0xbc, 0xc3, 0x8e, 0xeb, 0xbe, 0x24,
	0x39, 0xf5, 0x61, 0x5a, 0x6d, 0x65, 0xe6, 0xbf, 0x30, 0x07, 0x50, 0x5a, 0x18, 0x95, 0x00, 0x7d,
	0x5f, 0x6d, 0xe1, 0x00, 0x37, 0xec, 0x6e, 0x08, 0xe9, 0xad, 0xca, 0x55, 0x39, 0xf5, 0x51, 0x7d,
	0xc5, 0xf6, 0x2f, 0x09, 0x43, 0x23, 0xac, 0x5e, 0x10, 0x9e, 0xfa, 0x8f, 0xcf, 0xc7, 0xd0, 0x6a,
	0xc4, 0xd0, 0xbc, 0x5f, 0xfb, 0x93, 0xfb, 0xf1, 0xa2, 0xf2, 0x29, 0xe0, 0x5a, 0xcf, 0x6a, 0x4c,
	0x1a, 0xc2, 0x4d, 0x26, 0xbd, 0xec, 0xbf, 0x78, 0x04, 0x2a, 0xa6, 0x0f, 0x3a, 0x55, 0x97, 0xeb,
	0x57, 0x3a, 0xbf, 0x54, 0x93, 0xf8, 0x36, 0x6b, 0x57, 0x4f, 0x06, 0x97, 0x98, 0x6e, 0x5d, 0x84,
	0x4a, 0xd7, 0x05, 0x12, 0x76, 0x2d, 0xb3, 0x05, 0x78, 0x73, 0x0e, 0xe0, 0x20, 0x30, 0x47, 0x3b,
	0x0a, 0x8c, 0xcf, 0xcd, 0x06, 0xf3, 0xbf, 0x22, 0xd6, 0x13, 0x70, 0x35, 0x54, 0x93, 0x5c, 0xc8,
	0xe5, 0x68, 0x75, 0x63, 0x11, 0xd6, 0xde, 0x6b, 0xeb, 0xb3, 0xf7, 0x6a, 0x57, 0x67, 0xb0, 0x0a,
	0x0e, 0x09, 0x60, 0xc8, 0xb0, 0x2a, 0x94, 0x01, 0xef, 0xce, 0xa3, 0x6a, 0xba, 0xe9, 0xb8, 0x2e,
	0xe2, 0xa6, 0x1b, 0xca, 0x3d, 0x3e, 0xb8, 0x1d, 0x6f, 0x83, 0x9e, 0xdb, 0x6d, 0xd6, 0xbe, 0x04,
	0xa0, 0xf1, 0xa4, 0x2d, 0x70, 0x89, 0xdd, 0x26, 0x87, 0xe5, 0x9b, 0x15, 0x24, 0x60, 0x68, 0x34,
	0xe9, 0x89, 0x4a, 0xe0, 0x77, 0x5d, 0x63, 0xa0, 0xd9, 0xa4, 0x2b, 0x2a, 0x01, 0x7f, 0xca, 0x0e,
	0x5c, 0x17, 0xde, 0xc4, 0xb9, 0xb9, 0x79, 0x54, 0xbb, 0x39, 0x1f, 0x93, 0x9e, 0x36, 0xf6, 0x8d,
	0x31, 0x6f, 0xae, 0x21, 0xb7, 0x38, 0x11, 0x61, 0x53, 0x99, 0xeb, 0x74, 0x91, 0x81, 0x57, 0xae,
	0x49, 0x90, 0x5c, 0xab, 0xfd, 0xae, 0x23, 0x67, 0x83, 0xd1, 0x07, 0x18, 0xa3, 0x43, 0x76, 0x1d,
	0xe0, 0xff, 0x63, 0x9d, 0xf7, 0xb9, 0x1d, 0x3c, 0x47, 0xaa, 0x53, 0x69, 0x65, 0xf8, 0x91, 0x70,
	0xcd, 0xff, 0x8e, 0xa8, 0xd2, 0x5c, 0x79, 0xd5, 0xba, 0x33, 0x4d, 0x2f, 0x48, 0x0c, 0xbd, 0xca,
	0xc8, 0x4f, 0x2f, 0x41, 0x80, 0xa6, 0xf0, 0x07, 0xf6, 0xed, 0x99, 0xd6, 0x5f, 0xd4, 0xf6, 0x42,
	0x1b, 0xed, 0x7c, 0xd6, 0x46, 0xb7, 0x37, 0x6d, 0xf4, 0x21, 0x63, 0xc5, 0x62, 0x3c, 0x83, 0x75,
	0x21, 0x55, 0xa0, 0xb8, 0x26, 0xa1, 0x32, 0x53, 0x2b, 0xf7, 0x4d, 0xec, 0xd1, 0x3d, 0x36, 0xb8,
	0x56, 0x11, 0x3d, 0x77, 0x17, 0x87, 0xf8, 0x77, 0xc8, 0xf7, 0x95, 0xff, 0xaf, 0xe8, 0x07, 0xc2,
	0xdf, 0x5d, 0xd9, 0xa9, 0x5e, 0x58, 0xdf, 0xd3, 0xfc, 0xd8, 0xf4, 0x89, 0xf4, 0xe5, 0xa3, 0x5f,
	0x1e, 0x4c, 0x94, 0x9d, 0x2e, 0xc6, 0x27, 0x89, 0x9e, 0x3f, 0x1b, 0x0c, 0x92, 0xfc, 0x19, 0x0d,
	0xf9, 0x83, 0xc1, 0x33, 0x9a, 0x45, 0xc6, 0xdb, 0x34, 0xce, 0x0f, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x59, 0x52, 0xd8, 0x29, 0x0c, 0x00, 0x00,
}
