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

export class WithdrawerService {
  /**
   *
   */
  withdrawerServiceExportExchangeAccounts(
    params: {
      /**  */
      body: ExportExchangeAccountsReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ExportExchangeAccountsRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdraw/export';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceCreateWithdrawer(
    params: {
      /**  */
      body: CreateWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/add';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceDeleteWithdrawer(
    params: {
      /**  */
      body: DeleteWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DeleteWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceGetExchangeDepositOptions(
    params: {
      /**  */
      body: GetExchangeDepositOptionsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetExchangeDepositOptionsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/exchange/deposit/options';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceGetExchangeWithdrawOptions(
    params: {
      /**  */
      body: GetExchangeWithdrawOptionsRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetExchangeWithdrawOptionsResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/exchange/options';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceGetWithdrawer(
    params: {
      /**  */
      body: GetWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceListWithdrawer(
    params: {
      /**  */
      body: ListWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceCreateSubWithdrawer(
    params: {
      /**  */
      body: CreateOkexWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateOkexWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/sub/add';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceOkexDepositAddrAttach(
    params: {
      /**  */
      body: OkexDepositAddrAttachRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<OkexDepositAddrAttachResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/sub/deposit/attach';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceOkexDepositAddrDetach(
    params: {
      /**  */
      body: OkexDepositAddrDetachRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<OkexDepositAddrDetachResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/sub/deposit/detach';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceListDepositAddresses(
    params: {
      /**  */
      body: ListDepositAddressesRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListDepositAddressesResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/sub/deposit/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceListSubWithdrawer(
    params: {
      /**  */
      body: ListSubWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListSubWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/sub/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceUpdateWithdrawer(
    params: {
      /**  */
      body: UpdateWithdrawerRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UpdateWithdrawerResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceWithdraw(
    params: {
      /**  */
      body: WithdrawReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<WithdrawRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/withdraw';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  withdrawerServiceWithdrawStatus(
    params: {
      /**  */
      body: WithdrawStatusReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<WithdrawStatusRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/withdrawer/withdraw/status';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface CreateOkexWithdrawerRequest {
  /**  */
  parentId: string;

  /**  */
  label: string;

  /**  */
  secretKey: string;

  /**  */
  apiKey: string;

  /**  */
  proxy?: string;
}

export interface CreateOkexWithdrawerResponse {
  /**  */
  error?: string;
}

export interface CreateWithdrawerRequest {
  /**  */
  exchangeType: ExchangeType;

  /**  */
  label: string;

  /**  */
  secretKey: string;

  /**  */
  apiKey: string;

  /**  */
  proxy?: string;
}

export interface CreateWithdrawerResponse {
  /**  */
  withdrawer: Withdrawer;
}

export interface DeleteWithdrawerRequest {
  /**  */
  id: string;
}

export interface DeleteWithdrawerResponse {}

export interface DepositAddresses {
  /**  */
  addr: string;

  /**  */
  profileId?: string;

  /**  */
  tag?: string;

  /**  */
  profileLabel?: string;
}

export interface ExchangeWithdrawNetwork {
  /**  */
  network: string;

  /**  */
  minAmount: string;

  /**  */
  maxAmount: string;

  /**  */
  fee: string;
}

export interface ExchangeWithdrawOptions {
  /**  */
  token: string;

  /**  */
  amount: string;

  /**  */
  networks: ExchangeWithdrawNetwork[];
}

export interface ExportExchangeAccountsReq {}

export interface ExportExchangeAccountsRes {
  /**  */
  text: string;
}

export interface GetExchangeDepositOptionsRequest {
  /**  */
  withdrawerId: string;

  /**  */
  token: string;

  /**  */
  network: string;
}

export interface GetExchangeDepositOptionsResponse {
  /**  */
  addr: string;
}

export interface GetExchangeWithdrawOptionsRequest {
  /**  */
  withdrawerId: string;
}

export interface GetExchangeWithdrawOptionsResponse {
  /**  */
  tokens: ExchangeWithdrawOptions[];
}

export interface GetWithdrawerRequest {
  /**  */
  withdrawerId: string;
}

export interface GetWithdrawerResponse {
  /**  */
  withdrawer: Withdrawer;

  /**  */
  error?: string;
}

export interface ListDepositAddressesRequest {
  /**  */
  withdrawerId: string;
}

export interface ListDepositAddressesResponse {
  /**  */
  items: DepositAddresses[];
}

export interface ListSubWithdrawerRequest {
  /**  */
  withdrawerId: string;
}

export interface ListSubWithdrawerResponse {
  /**  */
  withdrawers: Withdrawer[];
}

export interface ListWithdrawerRequest {}

export interface ListWithdrawerResponse {
  /**  */
  withdrawers: Withdrawer[];
}

export interface OkexDepositAddrAttachRequest {
  /**  */
  profileId: string;

  /**  */
  okexDepositAddr: string;

  /**  */
  withdrawerId: string;
}

export interface OkexDepositAddrAttachResponse {}

export interface OkexDepositAddrDetachRequest {
  /**  */
  profileId: string;

  /**  */
  okexDepositAddr: string;

  /**  */
  withdrawerId: string;
}

export interface OkexDepositAddrDetachResponse {}

export interface UpdateWithdrawerRequest {
  /**  */
  withdrawerId: string;

  /**  */
  proxy: string;

  /**  */
  label: string;
}

export interface UpdateWithdrawerResponse {
  /**  */
  error?: string;
}

export interface WithdrawReq {
  /**  */
  network: string;

  /**  */
  token: string;

  /**  */
  amount: string;

  /**  */
  withdrawerId: string;

  /**  */
  profileId: string;
}

export interface WithdrawRes {
  /**  */
  withdrawId: string;

  /**  */
  errorMessage?: string;
}

export interface WithdrawStatusReq {
  /**  */
  withdrawId: string;
}

export interface WithdrawStatusRes {
  /**  */
  status: string;
}

export interface Withdrawer {
  /**  */
  id: string;

  /**  */
  exchangeType: ExchangeType;

  /**  */
  label: string;

  /**  */
  proxy: string;

  /**  */
  createdAt: Date;
}

export enum ExchangeType {
  'Binance' = 'Binance',
  'Okex' = 'Okex'
}
