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

export class ProfileService {
  /**
   *
   */
  profileServiceCreateProfile(
    params: {
      /**  */
      body: CreateProfileRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<CreateProfileResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/add';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceGetBalance(
    params: {
      /**  */
      body: GetBalanceRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetBalanceResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/balance';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceDeleteProfile(
    params: {
      /**  */
      body: DeleteProfileRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<DeleteProfileResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/delete';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceExportProfiles(
    params: {
      /**  */
      body: ExportProfilesReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ExportProfilesRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/export';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceGenerateProfiles(
    params: {
      /**  */
      body: GenerateProfilesReq;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GenerateProfilesRes> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/generate';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceGetProfile(
    params: {
      /**  */
      body: GetProfileRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<GetProfileResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/get';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceValidateLabel(
    params: {
      /**  */
      body: ValidateLabelRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ValidateLabelResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/label/validate';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceListProfile(
    params: {
      /**  */
      body: ListProfileRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<ListProfileResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/list';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceSearchProfilesNotConnectedToOkexDeposit(
    params: {
      /**  */
      body: SearchProfilesNotConnectedToOkexDepositRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SearchProfilesNotConnectedToOkexDepositResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/not_connected_okex/search';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceSearchProfile(
    params: {
      /**  */
      body: SearchProfileRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<SearchProfileResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/search';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
  /**
   *
   */
  profileServiceUpdateProfile(
    params: {
      /**  */
      body: UpdateProfileRequest;
    } = {} as any,
    options: IRequestOptions = {}
  ): Promise<UpdateProfileResponse> {
    return new Promise((resolve, reject) => {
      let url = basePath + '/api/gw/v1/profile/update';

      const configs: IRequestConfig = getConfigs('post', 'application/json', url, options);

      let data = params['body'];

      configs.data = data;

      axios(configs, resolve, reject);
    });
  }
}

export interface Balance {
  /**  */
  token: Token;

  /**  */
  amount: string;

  /**  */
  wei: string;

  /**  */
  error?: string;
}

export interface CreateProfileRequest {
  /**  */
  label: string;

  /**  */
  proxy?: string;

  /**  */
  mmskPk: string;

  /**  */
  meta?: string;

  /**  */
  okexAccount?: OkexAccount;
}

export interface CreateProfileResponse {
  /**  */
  profile: Profile;
}

export interface DeleteProfileRequest {
  /**  */
  id: string;
}

export interface DeleteProfileResponse {}

export interface ExportProfilesReq {}

export interface ExportProfilesRes {
  /**  */
  text: string;
}

export interface GenerateProfilesReq {
  /**  */
  count: string;
}

export interface GenerateProfilesRes {
  /**  */
  text: string;
}

export interface GetBalanceRequest {
  /**  */
  profileId: string;

  /**  */
  network: Network;
}

export interface GetBalanceResponse {
  /**  */
  balance: Balance[];
}

export interface GetProfileRequest {
  /**  */
  profileId: string;
}

export interface GetProfileResponse {
  /**  */
  profile: Profile;
}

export interface ListProfileRequest {}

export interface ListProfileResponse {
  /**  */
  profiles: Profile[];
}

export interface OkexAccount {
  /**  */
  subAccName: string;

  /**  */
  depositAddr: string;
}

export interface Profile {
  /**  */
  id: string;

  /**  */
  label: string;

  /**  */
  proxy?: string;

  /**  */
  mmskId: string;

  /**  */
  meta: string;

  /**  */
  createdAt: Date;

  /**  */
  okexAccount?: OkexAccount;

  /**  */
  userAgent: string;

  /**  */
  num: string;

  /**  */
  deletedAt?: Date;
}

export interface SearchProfileRequest {
  /**  */
  pattern: string;
}

export interface SearchProfileResponse {
  /**  */
  profiles: Profile[];
}

export interface SearchProfilesNotConnectedToOkexDepositRequest {}

export interface SearchProfilesNotConnectedToOkexDepositResponse {
  /**  */
  profiles: Profile[];
}

export interface UpdateProfileRequest {
  /**  */
  profileId: string;

  /**  */
  proxy?: string;

  /**  */
  meta?: string;

  /**  */
  okexAccount?: OkexAccount;

  /**  */
  label: string;

  /**  */
  userAgent: string;
}

export interface UpdateProfileResponse {
  /**  */
  profile: Profile;
}

export interface ValidateLabelRequest {
  /**  */
  label: string;

  /**  */
  profileId?: string;
}

export interface ValidateLabelResponse {
  /**  */
  valid: boolean;
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

export enum Token {
  'USDT' = 'USDT',
  'ETH' = 'ETH',
  'USDC' = 'USDC',
  'STG' = 'STG',
  'BNB' = 'BNB',
  'MATIC' = 'MATIC',
  'AVAX' = 'AVAX',
  'veSTG' = 'veSTG',
  'WETH' = 'WETH',
  'LUSD' = 'LUSD',
  'LSD' = 'LSD',
  'MUTE' = 'MUTE',
  'MAV' = 'MAV',
  'SPACE' = 'SPACE',
  'VC' = 'VC',
  'IZI' = 'IZI'
}
