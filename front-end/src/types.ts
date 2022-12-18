export interface Group {
  id: number;
  name: string;
  image: string;
  members: string[];
  creationDate: number;
  firstAlbum: string;
}

export interface GroupModel {
  id: number;
  name: string;
  image: string;
  creationDate: number;
}

export interface Concerts {
  id: number;
  datesLocations: {
    [name: string]: [dates: string[]];
  };
}
