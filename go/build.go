package main;

import ("net/http"
        "fmt"
        "io/ioutil"
        "strings"
        "os"
        "io"
        "runtime")


func main(){
    content := fetch();
    pretty := format(content);
    toUser := rmDupes(pretty);
    write(toUser);
}

func fetch()[]string{
    
    var server string = "https://raw.githubusercontent.com/roypur/hosts/master/src";
    
    resp,_ := http.Get(server);
    defer resp.Body.Close();
    body,_ := ioutil.ReadAll(resp.Body);
    
    src := strings.Split(strings.TrimSpace(string(body)), "\n");
    
    requestData := []string{};
    
    var j int = len(src) - 1;
    
    for i := 0; i<=j; i++ {
        resp,_ := http.Get(src[i]);
        defer resp.Body.Close();
        body,_ := ioutil.ReadAll(resp.Body);
        
        requestData = append(requestData, string(body));
    }
    
    
    var appendString string = strings.Join(requestData, "\n");
    
    retVal := strings.Split(appendString, "\n");
    
    return retVal;
}

func format(all []string)[]string{
    
    clean := []string{};
    
    for _,value := range all {
    
        value = strings.TrimSpace(value);
        
        var isComment bool = strings.HasPrefix(value, "#");
        var isFour bool = strings.HasPrefix(value, "127.0.0.1 ");
        var isSix bool = strings.HasPrefix(value, "::1 ");
    
        if(!isComment){
        
            value = strings.Split(value, "#")[0];
            value = strings.TrimSpace(value);
            value = strings.Replace(value, "\t", " ", -1);
            
            var old string = "";
            
            for old != value {
                old = value;
                value = strings.Replace(value, "  ", " ", 1);
            }
            
            clean = append(clean, value);

            if(isFour){
                clean = append(clean, strings.Replace(value, "127.0.0.1", "::1", 1)); //if ip is v4 add v6
            }else if(isSix){
                clean = append(clean, strings.Replace(value, "::1", "127.0.0.1", 1)); //if ip is v6 add v4
            }
        }
    }
    
    return clean;
}

func rmDupes(pretty []string)string{

    final := []string{};
    
    for _,value := range pretty {
    
        if(exists(value, final) == false){
            final = append(final, value);
        }
    }
    
    var retVal string = strings.Join(final, "\n");
    retVal = strings.TrimSpace(retVal);
    return retVal;
}


func exists(str string, list []string) bool{
    
    for _,v := range list{
    
        if(str == v){
            return true;
        }
    }
    return false;
}

func write(file string){
    
    var filename = "";
    
    if(runtime.GOOS == "linux"){
        filename = "/etc/hosts";
    }else if(runtime.GOOS == "windows"){
        filename = "C:/Windows/system32/drivers/etc/hosts";
    }
    
    f, err := os.Create(filename)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(" Write to file : " + filename)
    
    n, err := io.WriteString(f, file)

    if err != nil {
        fmt.Println(n, err)
    }

    f.Close()
}
