from HTMLParser import HTMLParser
from htmlentitydefs import name2codepoint
import glob, os

'''
web example: 

def show(self, req, resp_obj, id): 

<p id="t41" class="pln">    <span class="key">def</span> <span class="nam">show</span><span class="op">(</span><span class="nam">self</span><span class="op">,</span> <span class="nam">req</span><span class="op">,</span> <span class="nam">resp_obj</span><span class="op">,</span> <span class="nam">id</span><span class="op">)</span><span class="op">:</span><span class="strut">&nbsp;</span></p>

context = req.environ['nova.context']

<p id="t42" class="stm mis">        <span class="nam">context</span> <span class="op">=</span> <span class="nam">req</span><span class="op">.</span><span class="nam">environ</span><span class="op">[</span><span class="str">'nova.context'</span><span class="op">]</span><span class="strut">&nbsp;</span></p>
'''
#global variable
g_print = "" # print strings
g_now_tag = ""
g_now_id = ""
g_now_class = ""
g_now_function = ""
g_now_line_number = 0
g_pre_id = ""
g_num_tab_space_before_data = 0
g_num_tab_space_before_def = 0
g_num_insert_lines = 0
g_num_def_start = 0
g_num_def_end = 0
g_bool_find_def = False
g_bool_find_function = False
g_bool_find_t_id = False
g_bool_find_miss_statement = False
g_bool_find_source_file_path = False
g_bool_find_import = False
g_bool_find_end = False
g_bool_problem_file = False
g_chek_two_line_end = False
g_write_text = ""
g_touch_path = "/home/ubuntu/kentzu/temp1020/"

mis_num =""

class MyHTMLParser(HTMLParser):
    
    def handle_starttag(self, tag, attrs):
        global g_print,g_now_tag,g_now_id,g_now_class,g_now_function,g_num_tab_space_before_def,g_num_insert_lines,g_bool_find_source_file_path
        
        #print "Start tag:", tag
        g_now_tag = tag

        if tag == "b":
            g_bool_find_source_file_path = True
            print "tag ==b=",tag
        for attr in attrs:
            if ( "id" in attr): # ex attr = (id, t10) => right td id in the table.
                g_now_id = attr
            
            if ( "class" in attr):
                g_now_class = attr


    def handle_endtag(self, tag):
        # find html end tag in a line ex: </p>
        global g_print
        #print "End tag  :", tag

    def handle_data(self, data):
        # find html data in a line. ex: <>if a is not b:</>
        global g_print,g_now_tag,g_now_id,g_now_class,g_now_function,g_num_tab_space_before_def,g_num_insert_lines,g_num_def_start,g_num_def_end,g_bool_find_def,g_bool_find_function,g_bool_find_t_id,g_bool_find_miss_statement,g_num_tab_space_before_data,g_now_line_number,g_pre_id,g_chek_two_line_end,g_bool_find_import,g_write_text,g_bool_find_source_file_path,g_bool_find_end,g_touch_path,mis_num 
        
        
        # find source file path 
        if (g_bool_find_source_file_path):
            g_write_text  += "filepath::"+ data +" \n" 
            g_bool_find_source_file_path = False

        # start search in id=tnumber    
        if ( "id" in g_now_id and  "t" in g_now_id[1] and (g_now_id[1].replace("t", "")).isdigit() ):
            g_bool_find_t_id = True
            
            if( g_pre_id != g_now_id):
                
                # tab space number of data
                space = data.split(" ")
                g_num_tab_space_before_data = len(space) - 1

                
                # get  line number
                line_number = g_now_id[1].replace("t", "") # get line number
                g_pre_id = g_now_id
                if(line_number.isdigit()):
                    g_now_line_number = int(line_number)

                #find a new func or class
                if (g_bool_find_miss_statement ):
                    if(g_num_tab_space_before_def != 0 and g_num_tab_space_before_data <= g_num_tab_space_before_def and  "pln" not in g_now_class):
                    #     print "*==11111111111111"
                    #     print g_chek_two_line_end
                    #     print g_num_def_end
                    #     if(not g_chek_two_line_end):
                    #         g_chek_two_line_end = True
                    #     else:
                    #         g_bool_find_end = True
                    #         g_num_def_end = int(g_now_line_number) - 1
                    # else:
                    #     g_chek_two_line_end = False # check the end line 
            
        else:
            g_bool_find_t_id = False
        
        
        
        
        if (g_bool_find_t_id): # start to read line at id=tnumber
             
            # find number of import line
            if(not g_bool_find_import and "key" in g_now_class and (data == "from" or data == "import")):
                g_write_text  += str(g_now_line_number) + "::0::import os\n" 
                g_num_insert_lines += 1
                g_bool_find_import = True
                
            if (g_bool_find_miss_statement and g_bool_find_end and g_num_def_start> 0 and g_num_def_end> 0):
                if(g_num_def_end > g_num_def_start):
    
                    # the last line of function
                   
                    print "*ffffffffffff*g_now_line_number",  str(g_now_line_number)
                    print "**g_num_def_start", g_num_def_start
                    print "**g_num_def_end", g_num_def_end
                    print "**g_now_function", g_now_function
                    # set start and end line + has inserted lines
                    g_num_def_start += int(g_num_insert_lines)
                    g_num_def_end += int(g_num_insert_lines)
                    g_num_insert_lines += 2
                    code_space = g_num_tab_space_before_def + 4
                    """
                    print "*==g_num_insert_lines", g_num_insert_lines
                    print "*==g_num_def_start", g_num_def_start
                    print "*==g_num_def_end", g_num_def_end
                    print "*==g_now_function", g_now_function
                    """
                    #write to text
                    g_write_text  += str(g_num_def_start)+ "::"+str(code_space)+"::os.system(\"touch "+ g_touch_path+ file_title + "_"+g_now_function+" \")\n" 
                    g_write_text  += str(g_num_def_end)+"::"+str(code_space)+"::os.system(\"touch "+ g_touch_path + file_title + "_"+g_now_function+"_end \")\n" 
                    print "**g_write_text\n", g_write_text
                    #reset
                    g_bool_find_function = False
                    g_bool_find_def = False
                    g_bool_find_t_id = False
                    g_bool_find_miss_statement = False
                    g_num_tab_space_before_def = 0
                    g_chek_two_line_end = False
                    g_num_def_end = 0
                    g_num_def_start = 0
                    g_bool_find_end = False
            
            
        
            # find a def tag in a line    
            if data == "return" and g_num_tab_space_before_data <= 8:
                g_num_def_end = g_now_line_number
                g_bool_find_end = True
            
            
            
            # get function name
            if (g_bool_find_def and  "nam"  in g_now_class  and g_bool_find_function==False):
                g_bool_find_function = True
                g_now_function = data
  
            # check if miss line in this function 
            if ("stm mis" in g_now_class):
                print "*==xxxxxxxxxxxxx", str(g_now_line_number)
                g_bool_find_miss_statement = True
            
            #print "**data", data
            # find a def tag in a line    
            if data == "def" and g_num_tab_space_before_data <= 4:
                #mis_num += str(g_now_line_number) +","+data+","
                g_num_tab_space_before_def = g_num_tab_space_before_data
                g_num_def_start = int(g_now_line_number)
                print "g_num_def_start", g_num_def_start
                """
                print "*==g_num_insert_lines", g_num_insert_lines
                print "*==g_num_def_start", g_num_def_start
                print "*==g_num_def_end", g_num_def_end
                print "*==g_now_function", g_now_function
                print "*==g_bool_find_miss_statement", g_bool_find_miss_statement
                print "*==g_bool_find_end", g_bool_find_end
                """
                g_bool_find_function = False
                g_bool_find_def = True
            
            
        
        
            
            
                
    def handle_comment(self, data):
        global g_print
        g_print += "Comment  :" + ''.join(data) + "\n"
        #print "Comment  :", data

    def handle_entityref(self, name):
        global g_print
        c = unichr(name2codepoint[name])
        g_print += "Named ent:" + ''.join(c) + "\n"
        #print "Named ent:", c

    def handle_charref(self, name):
        global g_print
        if name.startswith('x'):
            c = unichr(int(name[1:], 16))
        else:
            c = unichr(int(name))
        g_print += "Num ent  :" + ''.join(c) + "\n"
        #print "Num ent  :", c

    def handle_decl(self, data):
        global g_print
        g_print += "Decl     :" + ''.join(data) + "\n"
        #print "Decl     :", data

parser = MyHTMLParser()

paths = []

#paths.append('/home/kristen/python_parse/testcase/controller_test')
paths.append('/home/kristen/python_parse/testcase/controller_1020')

#paths.append('/home/kristen/python_parse/testcase/fake_3')
for index in range(len(paths)):
    path = paths[index]
    print "path     :", path
    save_path = "/home/kristen/python_parse/save/insert_list/"
    '''
    if index == 0 :
        g_collect_mis_content +=  "\n controll\n"
        g_collect_miss_func += "\n controll \n"
    else: 
         g_collect_mis_content +=  "\n compute\n"
         g_collect_miss_func += "\n compute \n"
    '''
         
    for filename in os.listdir(path):
        print "=========1"
        #save_file = ""
        g_write_text = ""
        g_num_insert_lines = 0
        g_bool_find_import = False
        g_bool_problem_file = False
        g_print = "" # print strings
        g_now_tag = ""
        g_now_id = ""
        g_now_class = ""
        g_now_function = ""
        g_now_line_number = 0
        g_pre_id = ""
        g_num_tab_space_before_data = 0
        g_num_tab_space_before_def = 0
        g_num_insert_lines = 0
        g_num_def_start = 0
        g_num_def_end = 0
        g_bool_find_def = False
        g_bool_find_function = False
        g_bool_find_t_id = False
        g_bool_find_miss_statement = False
        g_bool_find_source_file_path = False
        g_bool_find_import = False
        g_bool_find_end = False
        g_bool_problem_file = False
        g_chek_two_line_end = False
        g_write_text = ""
        '''
        if("cinderclient" in filename):
            save_file = "cinder"
        if("glanceclient" in filename):
            save_file = "glance"
        if("keystoneclient" in filename):
            save_file = "keystonet"
        if("neutronclient" in filename):
            save_file = "neutron"
        '''
        if("nova" in filename):
            print "=========2"
            #save_file = "./nova_insert_list.txt"
            file_title = filename.replace( "_usr_lib_python2_7_dist-packages_", "")
            file_title =  file_title.replace("_py.html", "")
            # read file
            f = open(path +"/" +filename, "r")
            parser.feed(f.read())
            f.close()
            count_lines = len(g_write_text.split("::")) #if text contents "file name" + "import os", then doesn't wirte
            # write insert string in the file
            print "=========3"
            if(count_lines >= 6):
                print "=========4"
                if(g_bool_problem_file):
                    print "=========5"
                    save_file = save_path+ file_title + "_PROBLEM"
                else:
                    print "=========6"
                    save_file = save_path+ file_title
                with open(save_file, "w") as myfile:
                    print "=========7"
                    if(g_bool_find_miss_statement and g_bool_find_end and g_num_def_start> 0 and g_num_def_end> 0):
                        code_space = g_num_tab_space_before_def + 4
                        #write to text
                        g_write_text  += str(g_num_def_start)+ "::"+str(code_space)+"::os.system(\"touch "+ g_touch_path+ file_title + "_"+g_now_function+" \")\n" 
                        g_write_text  += str(g_num_def_end)+"::"+str(code_space)+"::os.system(\"touch "+ g_touch_path + file_title + "_"+g_now_function+"_end \")\n" 
                        #reset
                        g_bool_find_function = False
                        g_bool_find_def = False
                        g_bool_find_t_id = False
                        g_bool_find_miss_statement = False
                        g_num_tab_space_before_def = 0
                        g_chek_two_line_end = False
                        g_num_def_end = 0
                        g_num_def_start = 0
                        g_bool_find_end = False
                        g_bool_find_miss_statement = False
                    myfile.write(g_write_text.encode('utf8'))
                    myfile.close()
                    
                """
                with open("./save/nova/"+save_file+".txt", "a") as myfile:
                    myfile.write(g_collect_mis_content.encode('utf8'))
                    myfile.close()
                """
                """
                with open("./save/nova/"+save_file+"_func.txt", "a") as myfile:
                    myfile.write(g_collect_miss_func.encode('utf8'))
                    myfile.close()
                g_collect_mis_content = ""
                g_collect_miss_func = ""
                """



