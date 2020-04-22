export class Rect {
  public static fromXY(x: number, y: number) {
    return new Rect(0, x, 0, y);
  }

  constructor(private x1: number, private x2: number, private y1: number, private y2:number) {}

  public getX1(): number {
    return this.x1;
  }

  public getX2(): number {
    return this.x2;
  }

  public getY1(): number {
    return this.y1;
  }

  public getY2(): number {
    return this.y2;
  }
}