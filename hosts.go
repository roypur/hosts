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

    src := make(map[int]string);
    
    src[0] = "http://someonewhocares.org/hosts/ipv6/hosts";
    src[1] = "http://"
    
    appendString := []string{""};
    
    for _,v := range src{
        resp,_ := http.Get(v);
        defer resp.Body.Close();
        fmt.Println("done");
        body,_ := ioutil.ReadAll(resp.Body);
    
        var i int = 0;
        appendString[i]= string(body);        
        i++;
        
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
        var isFour bool = strings.HasPrefix(value, "127.0.0.1");
    
        if(!exists(value, clean) && !isComment && isFour){
            clean = append(clean, value);
            clean = append(clean, strings.Replace(value, "127.0.0.1", "::1", 1));
        }
    }
    
    for _,value := range all {
        value = strings.TrimSpace(value);
        
        var isComment bool = strings.HasPrefix(value, "#");
        var isFour bool = strings.HasPrefix(value, "::1 ");
    
        if(!exists(value, clean) && !isComment && isFour){
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
