/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useOlloReserveV1() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.OlloReserveV1.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryGetDenoms = ( options: any) => {
    const key = { type: 'QueryGetDenoms',  };    
    return useQuery([key], () => {
      return  client.OlloReserveV1.query.queryGetDenoms().then( res => res.data );
    }, options);
  }
  
  const QueryGetDenom = (denom: string,  options: any) => {
    const key = { type: 'QueryGetDenom',  denom };    
    return useQuery([key], () => {
      const { denom } = key
      return  client.OlloReserveV1.query.queryGetDenom(denom).then( res => res.data );
    }, options);
  }
  
  const QueryGetDenomWhitelist = (denom: string,  options: any) => {
    const key = { type: 'QueryGetDenomWhitelist',  denom };    
    return useQuery([key], () => {
      const { denom } = key
      return  client.OlloReserveV1.query.queryGetDenomWhitelist(denom).then( res => res.data );
    }, options);
  }
  
  const QueryDenomsFromCreator = (creator: string,  options: any) => {
    const key = { type: 'QueryDenomsFromCreator',  creator };    
    return useQuery([key], () => {
      const { creator } = key
      return  client.OlloReserveV1.query.queryDenomsFromCreator(creator).then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryGetDenoms,QueryGetDenom,QueryGetDenomWhitelist,QueryDenomsFromCreator,
  }
}