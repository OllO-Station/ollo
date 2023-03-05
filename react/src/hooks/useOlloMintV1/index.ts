/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloMintV1() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloMintV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryInflation = ( options: any) => {
    const key = { type: 'QueryInflation',  };    
    return useQuery([key], () => {
      return  client.OlloMintV1.query.queryInflation().then( res => res.data );
    }, options);
  }
  
  const QueryAnnualProvisions = ( options: any) => {
    const key = { type: 'QueryAnnualProvisions',  };    
    return useQuery([key], () => {
      return  client.OlloMintV1.query.queryAnnualProvisions().then( res => res.data );
    }, options);
  }
  
  const QueryDistribution = ( options: any) => {
    const key = { type: 'QueryDistribution',  };    
    return useQuery([key], () => {
      return  client.OlloMintV1.query.queryDistribution().then( res => res.data );
    }, options);
  }
  
  const QueryLastBlockTime = ( options: any) => {
    const key = { type: 'QueryLastBlockTime',  };    
    return useQuery([key], () => {
      return  client.OlloMintV1.query.queryLastBlockTime().then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryInflation,QueryAnnualProvisions,QueryDistribution,QueryLastBlockTime,
  }
}