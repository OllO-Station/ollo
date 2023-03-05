/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useIntertx() {
  const client = useClient();
  const QueryInterchainAccount = (owner: string, connection_id: string,  options: any) => {
    const key = { type: 'QueryInterchainAccount',  owner,  connection_id };    
    return useQuery([key], () => {
      const { owner,  connection_id } = key
      return  client.Intertx.query.queryInterchainAccount(owner, connection_id).then( res => res.data );
    }, options);
  }
  
  return {QueryInterchainAccount,
  }
}