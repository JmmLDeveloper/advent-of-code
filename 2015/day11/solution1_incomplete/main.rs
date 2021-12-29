
fn third_condition( vec_str : & Vec<u8>) -> bool {
    let mut condition_count = 0;
    let mut last_ch : Option<u8> = None;


    for &ch in vec_str {
        if last_ch.is_none() || last_ch.unwrap() != ch {
            
            last_ch = Some(ch)
        } else {
            condition_count = condition_count + 1;
            if condition_count == 2 {
                return true;
            }
            last_ch = None;
        }
    } 

    false
}

fn second_condition( vec_str : & Vec<u8>) -> bool {
    const I_ASCII : u8 = 105;
    const L_ASCII : u8 = 111;
    const O_ASCII : u8 = 108;

    return !vec_str.contains(&I_ASCII) && !vec_str.contains(&L_ASCII) && !vec_str.contains(&O_ASCII);
}

fn first_condition(vec_str : & Vec<u8>) -> bool {

    let mut condition_count = 0;
    let mut last_ch : Option<u8> = None;


    for &ch in vec_str {
        if last_ch.is_none() || last_ch.unwrap() + 1 == ch {
            condition_count = condition_count + 1;
            if condition_count == 3 {
                return true;
            }
            last_ch = Some(ch);
        } else {
            condition_count = 1;
            last_ch = Some(ch);
        }
    } 

    false
}

fn inc_pass( vec_str : &mut Vec<u8> ) {

    const Z_ASCII : u8 = 122;
    const A_ASCII : u8 = 97;


    let mut idx  = vec_str.len() - 1;

    while idx > 0  {
        let ch = vec_str[idx];
        if ch != Z_ASCII {
            vec_str[idx] = vec_str[idx] + 1 ;
            return
        } else {
            vec_str[idx] = A_ASCII;
            idx = idx - 1;
        }
    }

    println!("fatal Error");
}

fn main() {
    let input = String::from("vzbxkghb");
    let mut vec_str = input.as_bytes().to_vec();

    loop {
        if first_condition(&vec_str) && second_condition(&vec_str) && third_condition(&vec_str) {
            break;
        }

        inc_pass(&mut vec_str)
    }
    
    let result =  String::from_utf8(vec_str).unwrap();

    println!("the solution to part 1 is : {}",result);


}