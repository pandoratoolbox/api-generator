export interface ILeadInstagram {
  language?: string;
  website?: string;
  username?: string;
  following?: number;
  likes_average?: number;
  gender?: string;
  profile_picture_src?: string;
  likes_min?: number;
  comments_min?: number;
  id?: number;
  external_id?: number;
  followers?: number;
  likes_max?: number;
  comments_max?: number;
  bio?: string;
  tags_published?: object;
  comments_average?: number;
}

export class LeadInstagram {
  language?: string;
  website?: string;
  username?: string;
  following?: number;
  likes_average?: number;
  gender?: string;
  profile_picture_src?: string;
  likes_min?: number;
  comments_min?: number;
  id?: number;
  external_id?: number;
  followers?: number;
  likes_max?: number;
  comments_max?: number;
  bio?: string;
  tags_published?: object;
  comments_average?: number;

  constructor(data: ILeadInstagram) {
    this.language = data.language;
    this.website = data.website;
    this.username = data.username;
    this.following = data.following;
    this.likes_average = data.likes_average;
    this.gender = data.gender;
    this.profile_picture_src = data.profile_picture_src;
    this.likes_min = data.likes_min;
    this.comments_min = data.comments_min;
    this.id = data.id;
    this.external_id = data.external_id;
    this.followers = data.followers;
    this.likes_max = data.likes_max;
    this.comments_max = data.comments_max;
    this.bio = data.bio;
    this.tags_published = data.tags_published;
    this.comments_average = data.comments_average;
  }
}
