package condition

type ConditionModel struct {
	ObjectID   string `bson:"_id"`
	LabelID    int    `bson:"label_id"`
	Label      string `bson:"label_name"`
	Desciption string `bson:"description"`
	Effect     string `bson:"effect"`
	Solution   string `bson:"solution"`
	Prevention string `bson:"prevention"`
}
