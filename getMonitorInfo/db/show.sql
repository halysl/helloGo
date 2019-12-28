CREATE TABLE `test` (
  `target` varchar(255) NOT NULL,
  `uptime` varchar(255) DEFAULT NULL,
  `nodeload1` varchar(255) DEFAULT NULL,
  `nodeload5` varchar(255) DEFAULT NULL,
  `nodeload15` varchar(255) DEFAULT NULL,
  `node_cpu_usage` varchar(255) DEFAULT NULL,
  `node_memory_total` varchar(255) DEFAULT NULL,
  `node_memory_avaibable` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`target`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
