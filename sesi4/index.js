class Hello {
  constructor() {
    this.hello1 = this.hello1.bind(this);
    this.hello2 = this.hello2.bind(this);
  }

  async hello1() {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve("Hello1");
      }, 1000);
    });
  }

  async hello2() {
    return new Promise((resolve, reject) => {
      setTimeout(() => {
        resolve("Hello2");
      }, 2000);
    });
  }

  async init() {
    let result = await Promise.race([this.hello1(), this.hello2()]);

    console.log("result", result);
  }
}

let hello = new Hello();
hello.init();
