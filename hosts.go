package main;

import ("net/http"
        "fmt"
        "io/ioutil"
        "strings")


func main(){
    content := fetch();
    fmt.Println(clear(content));
}

func fetch()string{
    
    var server string = "https://raw.githubusercontent.com/roypur/hosts/master/src";
    
    resp,_ := http.Get(server);
    defer resp.Body.Close();
    body,_ := ioutil.ReadAll(resp.Body);
    
    src := strings.Split(str, "\n");
    
    appendString := []string{};
    
    var j int = len(src) - 1;
    
    for i := 0; i<=j; i++ {
        resp,_ := http.Get(src[i]);
        defer resp.Body.Close();
        body,_ := ioutil.ReadAll(resp.Body);
        
        appendString = append(appendString, string(body));        
    }
    
    var total string = strings.Join(appendString, "\n");

    return total;
}

func clear(str string)string{
    
    all := strings.Split(str, "\n");
    
    clean := []string{};
    
    
    //add ipv6 localhost if missing
    for _,value := range all {
    
        value = strings.TrimSpace(value);
        
        var isComment bool = strings.HasPrefix(value, "#");
        var isFour bool = strings.HasPrefix(value, "127.0.0.1");
    
        if(!exists(value, clean) && !isComment && isFour){
            clean = append(clean, value);
            clean = append(clean, strings.Replace(value, "127.0.0.1", "::1", 1));
        }
    }
    
    //add ipv4 localhost if missing
    for _,value := range all {
        value = strings.TrimSpace(value);
        
        var isComment bool = strings.HasPrefix(value, "#");
        var isSix bool = strings.HasPrefix(value, "::1 ");
    
        if(!exists(value, clean) && !isComment && isSix){
            clean = append(clean, value);
            clean = append(clean, strings.Replace(value, "::1 ", "127.0.0.1", 1));
        }
    }
    
    var retVal string = strings.Join(clean, "\n");
    return retVal;
}

func exists(str string, list []string) bool{
    for _,v := range list{
    
        if v == str {
            return true;
        }
    }
    return false;
}
