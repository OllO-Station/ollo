/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloClaimV1() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloClaimV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryClaimRecord = (address: string,  options: any) => {
    const key = { type: 'QueryClaimRecord',  address };    
    return useQuery([key], () => {
      const { address } = key
      return  client.OlloClaimV1.query.queryClaimRecord(address).then( res => res.data );
    }, options);
  }
  
  const QueryClaimRecordAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryClaimRecordAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloClaimV1.query.queryClaimRecordAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryGoal = (goalID: string,  options: any) => {
    const key = { type: 'QueryGoal',  goalID };    
    return useQuery([key], () => {
      const { goalID } = key
      return  client.OlloClaimV1.query.queryGoal(goalID).then( res => res.data );
    }, options);
  }
  
  const QueryGoalAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryGoalAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloClaimV1.query.queryGoalAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryAirdropSupply = ( options: any) => {
    const key = { type: 'QueryAirdropSupply',  };    
    return useQuery([key], () => {
      return  client.OlloClaimV1.query.queryAirdropSupply().then( res => res.data );
    }, options);
  }
  
  const QueryInitialClaim = ( options: any) => {
    const key = { type: 'QueryInitialClaim',  };    
    return useQuery([key], () => {
      return  client.OlloClaimV1.query.queryInitialClaim().then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryClaimRecord,QueryClaimRecordAll,QueryGoal,QueryGoalAll,QueryAirdropSupply,QueryInitialClaim,
  }
}