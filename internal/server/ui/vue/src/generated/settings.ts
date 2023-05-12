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

export class SettingsService {
  /**
   *
   */
  settingsServiceGetSettings(
    params: {
      /**  */
      body: GetSettingsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetSettingsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/settings/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  settingsServiceGetNetworkByRpc(
    params: {
      /**  */
      body: GetNetworkByRPCRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetNetworkByRPCResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/settings/network/rpc/validate';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  settingsServiceReset(
    params: {
      /**  */
      body: ResetRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ResetResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/settings/reset';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  settingsServiceUpdateSettings(
    params: {
      /**  */
      body: UpdateSettingsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UpdateSettingsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/api/v1/settings/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface GetNetworkByRPCRequest {
  /**  */
  endpoint: string;

  /**  */
  network: Network;
}

export interface GetNetworkByRPCResponse {
  /**  */
  valid: boolean;

  /**  */
  id: string;

  /**  */
  suggestedGasPrice: string;

  /**  */
  error?: string;
}

export interface GetSettingsRequest {}

export interface GetSettingsResponse {
  /**  */
  settings: Settings;
}

export interface ResetRequest {}

export interface ResetResponse {
  /**  */
  settings: Settings;
}

export interface Settings {
  /**  */
  userId?: string;

  /**  */
  bnb: SettingsNetwork;

  /**  */
  optimism: SettingsNetwork;

  /**  */
  arbitrum: SettingsNetwork;

  /**  */
  etherium: SettingsNetwork;

  /**  */
  polygon: SettingsNetwork;

  /**  */
  avalanche: SettingsNetwork;
}

export interface SettingsNetwork {
  /**  */
  id: string;

  /**  */
  useLimitGas: boolean;

  /**  */
  rpcEndpoint: string;

  /**  */
  gasTotal?: string;
}

export interface UpdateSettingsRequest {
  /**  */
  settings: Settings;
}

export interface UpdateSettingsResponse {
  /**  */
  settings: Settings;
}

export enum Network {
  'ARBITRUM' = 'ARBITRUM',
  'OPTIMISM' = 'OPTIMISM',
  'BinanaceBNB' = 'BinanaceBNB',
  'Etherium' = 'Etherium',
  'POLIGON' = 'POLIGON',
  'AVALANCHE' = 'AVALANCHE'
}