# Reative
用go实现的一个可以异步编程的玩意


  将一个函数根据所需的数据分为几个任务比如当执行数据库查询之前需要实例化数据库连接，又要对数据进行处理可将数据库实例化，和数据处理分开建立两个任务，当数据库实例化完之后将实例存储在map中，当需要用到数据库时从map中查询，在查询时，如果对应的map[string]chann 为空当前任务将进入阻塞，直到有值时再进行
