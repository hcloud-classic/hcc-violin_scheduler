[![pipeline status](http://118.130.73.5:8100/iitp-sds/violin-scheduler/badges/master/pipeline.svg)](http://118.130.73.5:8100/iitp-sds/violin-scheduler/pipelines)

[![coverage report](http://118.130.73.5:8100/iitp-sds/violin-scheduler/badges/master/coverage.svg)](http://118.130.73.5:8100/iitp-sds/violin-scheduler/commits/master)

[![go report](http://118.130.73.5:8100/iitp-sds/hcloud-badge/raw/feature/dev/hcloud-badge_violin-scheduler.svg)](http://118.130.73.5:8100/iitp-sds/hcloud-badge/raw/feature/dev/goreport_violin-scheduler)



# violin-scheduler

## 개발 현황

```shell

```



노드 선택 스케줄러

### 노드 스케쥴러 알고리즘 이란?

### 알고리즘 순서도

![NodeSelectAlgo](assets/NodeSelectAlgo.svg)





### 알고리즘 예시 1

![Exam1](assets/Exam1.svg)



### 알고리즘 예시 2

![Exam2](assets/Exam2.svg)

## GraphQL

### 알고리즘 graphql 예시

#### 1. 노드만 선택시

```shell
mutation _ {
  schedule_nodes(server_uuid: "qweqweqweqwe", cpu: 1, memory: 4, user_uuid: "1234",number_of_nodes: 4) {
    server_uuid
    cpu
    memory
    user_uuid
    number_of_nodes
  }
}


// selected nodes uuid list return
//Json Type
mutation _ {
  schedule_nodes(server_uuid: "qweqweqweqwe", cpu: 0, memory: 0, nr_node: 2) {
    node_uuid
  }
}


//String type
mutation _ {
	selected_nodes (server_uuid: "qweqweqweqwe", cpu: 0, memory: 0, nr_node: 2)
}
```

