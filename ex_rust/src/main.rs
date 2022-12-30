use warp::Filter;
use std::process::Command;

#[tokio::main]
async fn web_server() {
    let hello = warp::get()
        .map(|| format!("Hello, world!"));
    warp::serve(hello)
        .run(([127, 0, 0, 1], 11337))
        .await;
    // println!("Hello, world!");
}

// struct Question {
//     id: QuestionId,
//     title: String,
//     content: String,
//     tags: Option<Vec<String>>,
// }
// struct QuestionId(String);

// impl Question {
//     fn new(
//         id: QuestionId,
//         title: String,
//         content: String,
//         tags: Option<Vec<String>>,
//     ) -> Self {
//         Question {
//             id,
//             title,
//             content,
//             tags,
//         }
//     }

//     // fn update_title(&self, new_title: String) -> Self {
//     //     Question::new(self.id, new_title, self.content, self.tags)
//     // }

// }

fn main() {
    // let question = Question::new(
    //     "1",
    //     "First Question",
    //     "This is the first question",
    //     ["faq"]
    // );
    // println!("{}", question);
    // let mut cmd = Command::new("dir");
    // let res = cmd.output();
    // assert!(res.is_ok());

    let y:u32 = "10".parse().unwrap();
    println!("{}", y);
    let c1 = 'a';
    let c2 = '5';
    let c3 = '\u{263A}';
    println!("c1 = {}, c2 = {}, c3 = {}", c1, c2, c3);

    let tup : (u8, char, f32) = (10, 'a', 10.5);
    let first = tup.2;
    println!("{}", first);

    let mut arr: [i32; 10];
    arr = [0; 10];
    arr[2] = 2;
    let num1 = arr[2];
    println!("num1 = {}", num1);

    println!("result = {}", loop_func(8));
}

fn loop_func(count : u8) -> u8 {
    let mut x = 0;
    let result = loop {
        x += 1;
        println!("x = {}", x);
        if x == count { break 2*x }
    };
    let mut n = result;
    while n > 0 {
        n /= 2;
        println!("n = {}", n);
    }

    let arr = [10, 20, 30, 40];
    for item in arr.iter() {
        println!("item = {}", item);
    }

    return result;
}