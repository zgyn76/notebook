service 怎么和 deploy 及 pod 产生关联？
service根据 spec -> selector 关联pod，不与deploy产生关联。