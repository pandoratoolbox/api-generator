import { ILead, Lead } from "./lead";
import { IOrder, Order } from "./order";

export interface IDataPlatform {
  id?: number;
  name?: string;
  elasticsearch_index?: string;
  lead?: ILead;
  order?: IOrder[];
}

export class DataPlatform {
  id?: number;
  name?: string;
  elasticsearch_index?: string;
  lead?: Lead;
  order?: Order[];

  constructor(data: IDataPlatform) {
    this.id = data.id;
    this.name = data.name;
    this.elasticsearch_index = data.elasticsearch_index;
    this.lead = data.lead ? new Lead(data.lead) : undefined;
    this.order = data.order?.map((i) => {
      return new Order(i);
    });
  }
}
