import { IDataPlatform, DataPlatform } from "./data_platform";

export interface ILead {
  updated_at?: number | string | Date;
  external_id?: string;
  data_platform?: IDataPlatform;
  id?: number;
  data_platform_id?: number;
}

export class Lead {
  updated_at?: Date;
  external_id?: string;
  data_platform?: DataPlatform;
  id?: number;
  data_platform_id?: number;

  constructor(data: ILead) {
    this.updated_at = data.updated_at ? new Date(data.updated_at) : undefined;
    this.external_id = data.external_id;
    this.data_platform = data.data_platform
      ? new DataPlatform(data.data_platform)
      : undefined;
    this.id = data.id;
    this.data_platform_id = data.data_platform_id;
  }
}
