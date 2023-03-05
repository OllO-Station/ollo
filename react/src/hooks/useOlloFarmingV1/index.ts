/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloFarmingV1() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloFarmingV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryPlans = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryPlans', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.OlloFarmingV1.query.queryPlans(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryPlan = (plan_id: string,  options: any) => {
    const key = { type: 'QueryPlan',  plan_id };    
    return useQuery([key], () => {
      const { plan_id } = key
      return  client.OlloFarmingV1.query.queryPlan(plan_id).then( res => res.data );
    }, options);
  }
  
  const QueryStakings = (farmer: string, query: any, options: any) => {
    const key = { type: 'QueryStakings',  farmer, query };    
    return useQuery([key], () => {
      const { farmer,query } = key
      return  client.OlloFarmingV1.query.queryStakings(farmer, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryTotalStakings = (staking_coin_denom: string,  options: any) => {
    const key = { type: 'QueryTotalStakings',  staking_coin_denom };    
    return useQuery([key], () => {
      const { staking_coin_denom } = key
      return  client.OlloFarmingV1.query.queryTotalStakings(staking_coin_denom).then( res => res.data );
    }, options);
  }
  
  const QueryRewards = (farmer: string, query: any, options: any) => {
    const key = { type: 'QueryRewards',  farmer, query };    
    return useQuery([key], () => {
      const { farmer,query } = key
      return  client.OlloFarmingV1.query.queryRewards(farmer, query ?? undefined).then( res => res.data );
    }, options);
  }
  
  const QueryCurrentEpochDays = ( options: any) => {
    const key = { type: 'QueryCurrentEpochDays',  };    
    return useQuery([key], () => {
      return  client.OlloFarmingV1.query.queryCurrentEpochDays().then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryPlans,QueryPlan,QueryStakings,QueryTotalStakings,QueryRewards,QueryCurrentEpochDays,
  }
}