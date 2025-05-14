MetaData DataType {
    Int64 Price : `价格，N13(4)`,
    Int64 Qty : `数量，N15(2)`
    Int64 Amt : `金额，N18(4)`
    Int64 SeqNum : `消息序号`
    uInt16 Boolean : `1=True/Yes,0=False/No`
    uInt32 Length : `长度表示字节为单位的数据长度，正数`
    Int64 LocalTimeStamp : `本地时间戳，YYYYMMDDHHMMSSsss（毫秒），YYYY = 0000-9999, MM = 01-12, DD = 01-31, HH = 00-23, MM = 00-59, SS = 00-60 (秒)，sss=000-999 (毫秒)。`
    uInt32 NumInGroup : `重复数，表示重复组的个数，正数`
    uInt32 LocalMktDate : `本地市场日期，格式 YYYYMMDD，YYYY =     0000-9999, MM = 01-12, DD = 01-31`
    char[8] SecurityID : `证券代码`
    char[6] PBUID : `交易单元代码`
    char[12] AccountID : `证券账户`
    char[4] BranchID : `营业部代码`
    char[6] MemberID : `交易商代码`
    char[2] InvestorType : `交易主体类型，01=自营，02=资管，03=机构经  纪，04=个人经纪`
    char[10] InvestorID : `交易主体代码`
    char[8] TraderCode : `交易员代码`
    char[120] InvestorName : `客户名称，可包含中文字符，最长 120 个字节`
}

packet Logon {
    SenderCompID : `发送方代码`,
    TargetCompID : `接收方代码`,
    HeartBtInt ：`心跳间隔，单位为秒。订单管理系统登陆时提供给交易网关`,
    Password ：`密码`,
    DefaultApplVerID ：`二进制协议版本，即通信版本号`,
}