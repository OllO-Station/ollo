/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloLiquidityV1() {
  const client = useClient();
  const QueryLiquidityPools = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryLiquidityPools', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloLiquidityV1.query.queryLiquidityPools(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryLiquidityPool = (pool_id: string,  options: any) => {
    const key = { type: 'QueryLiquidityPool',  pool_id };    
    return useQuery([key], () => {
      const { pool_id } = key
      return  client.OlloLiquidityV1.query.queryLiquidityPool(pool_id).then( res => res.data );
    }, options);
  }
  
  const QueryLiquidityPoolByPoolCoinDenom = (pool_coin_denom: string,  options: any) => {
    const key = { type: 'QueryLiquidityPoolByPoolCoinDenom',  pool_coin_denom };    
    return useQuery([key], () => {
      const { pool_coin_denom } = key
      return  client.OlloLiquidityV1.query.queryLiquidityPoolByPoolCoinDenom(pool_coin_denom).then( res => res.data );
    }, options);
  }
  
  const QueryLiquidityPoolByReserveAcc = (reserve_acc: string,  options: any) => {
    const key = { type: 'QueryLiquidityPoolByReserveAcc',  reserve_acc };    
    return useQuery([key], () => {
      const { reserve_acc } = key
      return  client.OlloLiquidityV1.query.queryLiquidityPoolByReserveAcc(reserve_acc).then( res => res.data );
    }, options);
  }
  
  const QueryLiquidityPoolBatch = (pool_id: string,  options: any) => {
    const key = { type: 'QueryLiquidityPoolBatch',  pool_id };    
    return useQuery([key], () => {
      const { pool_id } = key
      return  client.OlloLiquidityV1.query.queryLiquidityPoolBatch(pool_id).then( res => res.data );
    }, options);
  }
  
  const QueryPoolBatchSwapMsgs = (pool_id: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPoolBatchSwapMsgs',  pool_id, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { pool_id,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloLiquidityV1.query.queryPoolBatchSwapMsgs(pool_id, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPoolBatchSwapMsg = (pool_id: string, msg_index: string,  options: any) => {
    const key = { type: 'QueryPoolBatchSwapMsg',  pool_id,  msg_index };    
    return useQuery([key], () => {
      const { pool_id,  msg_index } = key
      return  client.OlloLiquidityV1.query.queryPoolBatchSwapMsg(pool_id, msg_index).then( res => res.data );
    }, options);
  }
  
  const QueryPoolBatchDepositMsgs = (pool_id: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPoolBatchDepositMsgs',  pool_id, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { pool_id,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloLiquidityV1.query.queryPoolBatchDepositMsgs(pool_id, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPoolBatchDepositMsg = (pool_id: string, msg_index: string,  options: any) => {
    const key = { type: 'QueryPoolBatchDepositMsg',  pool_id,  msg_index };    
    return useQuery([key], () => {
      const { pool_id,  msg_index } = key
      return  client.OlloLiquidityV1.query.queryPoolBatchDepositMsg(pool_id, msg_index).then( res => res.data );
    }, options);
  }
  
  const QueryPoolBatchWithdrawMsgs = (pool_id: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPoolBatchWithdrawMsgs',  pool_id, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { pool_id,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloLiquidityV1.query.queryPoolBatchWithdrawMsgs(pool_id, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPoolBatchWithdrawMsg = (pool_id: string, msg_index: string,  options: any) => {
    const key = { type: 'QueryPoolBatchWithdrawMsg',  pool_id,  msg_index };    
    return useQuery([key], () => {
      const { pool_id,  msg_index } = key
      return  client.OlloLiquidityV1.query.queryPoolBatchWithdrawMsg(pool_id, msg_index).then( res => res.data );
    }, options);
  }
  
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloLiquidityV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  return {QueryLiquidityPools,QueryLiquidityPool,QueryLiquidityPoolByPoolCoinDenom,QueryLiquidityPoolByReserveAcc,QueryLiquidityPoolBatch,QueryPoolBatchSwapMsgs,QueryPoolBatchSwapMsg,QueryPoolBatchDepositMsgs,QueryPoolBatchDepositMsg,QueryPoolBatchWithdrawMsgs,QueryPoolBatchWithdrawMsg,QueryParams,
  }
}