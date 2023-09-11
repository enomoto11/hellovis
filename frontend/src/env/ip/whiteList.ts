export class IP {
  private whiteList: string[];

  constructor() {
    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    const ip1 = process.env.REACT_APP_IP_ADDRESS!;
    this.whiteList = [ip1];
  }

  public getWhiteList(): string[] {
    return [...this.whiteList];
  }
}
