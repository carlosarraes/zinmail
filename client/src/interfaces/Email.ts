export interface SingleMail {
  from: string;
  to: string;
  date: string;
  subject: string;
  content: string;
}

export interface Email {
  id: string;
  from: string;
  to: string;
  content: string;
  subject: string;
  date: string;
  highlight: string[];
}
