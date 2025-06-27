options {
	string_pre_fix_len_type = u16;
	repeat_pre_fix_size_type = u16;
}

packet SampleBinary {
    uint16 MsgType `消息类型`,
    u16 BodyLenght  = lengthof(Body)   `消息体长度`,
    match MsgType  as  Body {
		1 : Logon,
		2 : Logout,
		3 : Heartbeat,
		4 : RiskControlRequest,
		5 : RiskControlResponse,
	}
}

packet Logon {
    char[10] UserName `用户名`,
    string Password `密码`,
    uint64 ClientId `客户端ID`,
    u16 HeartbeatInterval `心跳间隔`,
}

packet Logout {
    char[10] UserName `用户名`,
    uint64 ClientId `客户端ID`,
}

packet Heartbeat {
}

packet RiskControlRequest {
    string UniqueOrderId `唯一订单号`,
    char[16] ClOrdID `客户订单号`,
    char[3] MarketID `市场id`,
    char[12] SecurityID `证券代码`,
    char Side `买卖方向`,
    char OrderType `订单类型`,
    u64 Price `价格`,
    u32 Qty `数量`,
    repeat string ExtraInfo `附加信息`,
    repeat SubOrder {
        char[16] ClOrdID `子订单号`,
        u64 Price `子订单价格`,
        u32 Qty `子订单数量`,
    }
}

packet RiskControlResponse {
    string UniqueOrderId `唯一订单号`,
    i32 Status `状态`,
    string Msg `结果信息`,
    repeat Detail,
}

packet Detail {
		string RuleName `规则名称`,
		u16 Code `原因代码`,
	}
