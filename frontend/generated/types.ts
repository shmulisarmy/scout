export type ServerState[Live_Scout] = {
  state: Live_Scout;
  key: string;
  client_list: any[];
}

export type Live_Scout = {
  links: { [key: string]:  };
  notes: string;
  as_of: string;
  scouted: boolean;
  to_scout_for: string;
}

