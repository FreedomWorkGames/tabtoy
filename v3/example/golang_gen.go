// Generated by github.com/davyxu/tabtoy
// DO NOT EDIT!!
// Version: 3.0.0
package example


type TableType int32
const (	
	TableType_None = 0 //  
	TableType_Data = 1 // 数据表 
	TableType_Symbol = 2 // 类型表 
	TableType_KeyValue = 3 // 键值表 
)

var (
	TableTypeMapperValueByName = map[string]int32{ 
		"None": 0, //  
		"Data": 1, // 数据表 
		"Symbol": 2, // 类型表 
		"KeyValue": 3, // 键值表 
	}

	TableTypeMapperNameByValue = map[int32]string{ 
		 0: "None", //  
		 1: "Data", // 数据表 
		 2: "Symbol", // 类型表 
		 3: "KeyValue", // 键值表 
	}
)

type ActorType int32
const (	
	ActorType_None = 0 //  
	ActorType_Pharah = 1 // 法鸡 
	ActorType_Junkrat = 2 // 狂鼠 
	ActorType_Genji = 3 // 源氏 
	ActorType_Mercy = 4 // 天使 
)

var (
	ActorTypeMapperValueByName = map[string]int32{ 
		"None": 0, //  
		"Pharah": 1, // 法鸡 
		"Junkrat": 2, // 狂鼠 
		"Genji": 3, // 源氏 
		"Mercy": 4, // 天使 
	}

	ActorTypeMapperNameByValue = map[int32]string{ 
		 0: "None", //  
		 1: "Pharah", // 法鸡 
		 2: "Junkrat", // 狂鼠 
		 3: "Genji", // 源氏 
		 4: "Mercy", // 天使 
	}
)



type TableField struct{ 
	Kind string `tb_name:"种类"` 
	ObjectType string `tb_name:"对象类型"` 
	Name string `tb_name:"标识名"` 
	FieldName string `tb_name:"字段名"` 
	FieldType string `tb_name:"字段类型"` 
	Value string `tb_name:"值"` 
	IsArray bool `tb_name:"数组"` 
	Splitter string `tb_name:"切割符"` 
}

type FieldType struct{ 
	InputFieldName string `tb_name:"输入字段"` 
	GoFieldName string `tb_name:"Go字段"` 
	CSFieldName string `tb_name:"C#字段"` 
	DefaultValue string `tb_name:"默认值"` 
}

type TablePragma struct{ 
	TableType TableType `tb_name:"表类型"` 
	TableName string `tb_name:"表名"` 
	TableFileName string `tb_name:"表文件名"` 
	IsVertical bool `tb_name:"垂直表"` 
}

type ExampleData struct{ 
	ID int32 `tb_name:"任务ID"` 
	Name string `tb_name:"名称"` 
	Rate float32 `tb_name:"比例"` 
	Type ActorType `tb_name:"类型"` 
	Skill []int32 `tb_name:"技能列表"` 
}

type ExampleKV struct{ 
	ServerIP string `tb_name:"服务器IP"` 
	ServerPort uint16 `tb_name:"服务器端口"` 
	GroupID []int32 `tb_name:"分组"` 
}


// Combine struct
type Config struct { 
	ExampleData []*ExampleData // table: ExampleData 
	ExampleKV []*ExampleKV // table: ExampleKV 
}