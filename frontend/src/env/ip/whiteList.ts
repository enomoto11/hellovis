export class IP {
  private whiteList: string[];

  constructor() {
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const ip1 = process.env.REACT_APP_IP_ADDRESS!;
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const ip2 = process.env.REACT_APP_IP_ADDRESS2!;

    this.whiteList = [ip1,ip2];
  }

  public getWhiteList(): string[] {
    return [...this.whiteList];
  }
}
