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
    
    src := strings.Split(strings.TrimSpace(string(body)), "\n");
    
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

    for _,value := range all {
    
        value = strings.TrimSpace(value);
        
        var isComment bool = strings.HasPrefix(value, "#");
        var isFour bool = strings.HasPrefix(value, "127.0.0.1 ");
        var isSix bool = strings.HasPrefix(value, "::1 ");
    
        if(!exists(value, clean) && !isComment){
        
            value = strings.Split(value, "#")[0];
            value = strings.TrimSpace(value);
            value = strings.Replace(value, "\t", " ", -1);
            value = strings.Replace(value, "  ", " ", -1);
        
            clean = append(clean, value);

            if(isFour){
                clean = append(clean, strings.Replace(value, "127.0.0.1", "::1", 1)); //if ip is v4 add v6
            }else if(isSix){
                clean = append(clean, strings.Replace(value, "::1", "127.0.0.1", 1)); //if ip is v6 add v4
            }
        }
    }
    
    var retVal string = strings.Join(clean, "\n");
    return retVal;
}

func exists(str string, list []string) bool{
    var toSix string = strings.Replace(str, "127.0.0.1 ", "::1 ", 1);
    var toFour string = strings.Replace(str, "::1 ", "127.0.0.1 ", 1);
    
    for _,v := range list{
    
        if (toSix == v || toFour == v) {
            return true;
        }
    }
    return false;
}
