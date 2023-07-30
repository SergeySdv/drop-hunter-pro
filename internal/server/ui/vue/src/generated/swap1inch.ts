/** Generate by swagger-axios-codegen */
// @ts-nocheck
/* eslint-disable */

/** Generate by swagger-axios-codegen */
// @ts-nocheck
/* eslint-disable */
export interface IRequestOptions {
  headers?: any;
  /** only in axios interceptor config*/
  loading: boolean;
  showError: boolean;
}

export interface IRequestPromise<T = any> extends Promise<IRequestResponse<T>> {}

export interface IRequestResponse<T = any> {
  data: T;
  status: number;
  statusText: string;
  headers: any;
  config: any;
  request?: any;
}

export interface IRequestInstance {
  (config: any): IRequestPromise;
  (url: string, config?: any): IRequestPromise;
  request<T = any>(config: any): IRequestPromise<T>;
}

export interface IRequestConfig {
  method?: any;
  headers?: any;
  url?: any;
  data?: any;
  params?: any;
}

// Add options interface
export interface ServiceOptions {
  axios?: IRequestInstance;
  /** only in axios interceptor config*/
  loading: boolean;
  showError: boolean;
}

// Add default options
export const serviceOptions: ServiceOptions = {};

// Instance selector
export function axios(configs: IRequestConfig, resolve: (p: any) => void, reject: (p: any) => void): Promise<any> {
  if (serviceOptions.axios) {
    return serviceOptions.axios
      .request(configs)
      .then(res => {
        resolve(res.data);
      })
      .catch(err => {
        reject(err);
      });
  } else {
    throw new Error('please inject yourself instance like axios  ');
  }
}

export function getConfigs(method: string, contentType: string, url: string, options: any): IRequestConfig {
  const configs: IRequestConfig = {
    loading: serviceOptions.loading,
    showError: serviceOptions.loading,
    ...options,
    method,
    url
  };
  configs.headers = {
    ...options.headers,
    'Content-Type': contentType
  };
  return configs;
}

export const basePath = '';

export interface IList<T> extends Array<T> {}
export interface List<T> extends Array<T> {}
export interface IDictionary<TValue> {
  [key: string]: TValue;
}
export interface Dictionary<TValue> extends IDictionary<TValue> {}

export interface IListResult<T> {
  items?: T[];
}

export class ListResultDto<T> implements IListResult<T> {
  items?: T[];
}

export interface IPagedResult<T> extends IListResult<T> {
  totalCount?: number;
  items?: T[];
}

export class PagedResultDto<T = any> implements IPagedResult<T> {
  totalCount?: number;
  items?: T[];
}

// customer definition
// empty

export class Swap1InchService {
  /**
   *
   */
  swap1InchServiceGetNetworks(
    params: {
      /**  */
      body: GetNetworksRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetNetworksResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/swap1inch/networks';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  swap1InchServiceGetSwapOptions(
    params: {
      /**  */
      body: GetSwapOptionsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetSwapOptionsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/swap1inch/swap/options';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  swap1InchServiceGetTokens(
    params: {
      /**  */
      body: GetTokensRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetTokensResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/swap1inch/tokens';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface GetNetworksRequest {}

export interface GetNetworksResponse {
  /**  */
  networks: Network[];
}

export interface GetSwapOptionsProtocol {
  /**  */
  name: string;

  /**  */
  estimatedGas: string;
}

export interface GetSwapOptionsRequest {
  /**  */
  fromTokenAddr: string;

  /**  */
  toTokenAddr: string;

  /**  */
  amount: string;

  /**  */
  network: Network;
}

export interface GetSwapOptionsResponse {
  /**  */
  protocols: GetSwapOptionsProtocol[];
}

export interface GetTokensRequest {
  /**  */
  network: Network;
}

export interface GetTokensResponse {
  /**  */
  tokens: Swap1inchToken[];
}

export interface Swap1inchToken {
  /**  */
  code: string;

  /**  */
  name: string;

  /**  */
  addr: string;
}

export enum Network {
  'ARBITRUM' = 'ARBITRUM',
  'OPTIMISM' = 'OPTIMISM',
  'BinanaceBNB' = 'BinanaceBNB',
  'Etherium' = 'Etherium',
  'POLIGON' = 'POLIGON',
  'AVALANCHE' = 'AVALANCHE',
  'GOERLIETH' = 'GOERLIETH',
  'ZKSYNCERA' = 'ZKSYNCERA',
  'ZKSYNCERATESTNET' = 'ZKSYNCERATESTNET',
  'ZKSYNCLITE' = 'ZKSYNCLITE'
}
