import Ui_pam_dialog
import sys
from PyQt5.QtWidgets import QApplication, QDialog

class Mainpam(QDialog):

    def __init__(self, parent=None):
        super(QDialog, self).__init__(parent)
        self.ui = Ui_pam_dialog.Ui_Dialog()
        self.ui.setupUi(self)

    def passwd_conf(self):
        print("configue passwd")
        minlen = self.ui.minlen.value()
        lcredit = self.ui.lcredit.value()
        minclass = self.ui.minclass.value()
        ucredit = self.ui.ucredit.value()
        dcredit = self.ui.dcredit.value()
        ocredit = self.ui.ocredit.value()
        retry = self.ui.retry.value()
        remember = self.ui.remember.value()

        # minlen = 7 ; minclass=4 ; lcredit = 0 ; ucredit=0 ; dcredit = 0 ; ocredit=0 ; retry=4
        # remember = 8

        with open('/etc/pam.d/system-auth', mode='rb+') as f:
            while True:#首先删去已有配置行
                # try:
                line = f.readline()  # 逐行读取
                # except IndexError:  # 超出范围则退出
                #     break
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                if ((line_list[0] == 'password' and  line_list[1] == 'requisite'  and line_list[2] == 'pam_pwquality.so')
                    or (line_list[0] == 'password' and  line_list[1] == 'requisite'  and line_list[2] == 'pam_pwhistory.so' )):                   
                    #删除这一句
                    rest = f.read()  
                    f.seek(-len(rest), 1)
                    f.seek(-len(line), 1) 
                    f.truncate() 
                    f.write(rest)
                    f.seek(-len(rest), 1)
            f.seek(0)
            while True:
                try:
                    line = f.readline()  # 逐行读取
                except IndexError:  # 超出范围则退出
                    break
                line_str = line.decode().splitlines()
                line_list = line_str[0].split()
                if line_list == []:
                    continue
                if line_list[0] == 'password':                
                    f.seek(-len(line), 1)
                    rest = f.read() # rest保留这一句
                    f.seek(-len(rest), 1)
                    f.truncate()
                    content1 = 'password    requisite     pam_pwquality.so try_first_pass local_users_only \
minlen={} minclass={} lcredit={} ucredit={} dcredit={} ocredit={} retry={}\
\n'.format(minlen,minclass,lcredit,ucredit,dcredit,ocredit,retry)
                    content2 = 'password    requisite     pam_pwhistory.so remember={}\n'.format(remember)
                    f.write(content1.encode())
                    f.write(content2.encode())  
                    f.write(rest)
                    f.seek(-len(rest), 1)
                    break

        
        
    def usrlogin_conf(self):
        print("usrlogin---")

        user_deny = self.ui.user_deny.value()
        user_unlocktime = self.ui.user_unlocktime.value()
        with open('/etc/pam.d/system-auth', mode='rb+') as f:
            while True:#首先删去已有配置行
                # try:
                line = f.readline()  # 逐行读取
                # except IndexError:  # 超出范围则退出
                #     break
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                try:
                    if ((line_list[0] == 'auth' and  line_list[1] == 'required'  and line_list[2] == 'pam_faillock.so' and line_list[3] == 'preauth')
                        or (line_list[0] == 'auth' and  line_list[1] == '[default=die]'  and line_list[2] == 'pam_faillock.so' and line_list[3] == 'authfail' )
                        or (line_list[0] == 'auth' and  line_list[1] == 'sufficient'  and line_list[2] == 'pam_faillock.so' and line_list[3] == 'authsucc' )):                   
                        #删除这一句
                        rest = f.read()  
                        f.seek(-len(rest), 1)
                        f.seek(-len(line), 1) 
                        f.truncate() 
                        f.write(rest)
                        f.seek(-len(rest), 1)
                except IndexError:
                    break
            f.seek(0)
            while True:
                #try:
                line = f.readline()  # 逐行读取
                # except IndexError:  # 超出范围则退出
                #     break
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                if line_list[0] == 'auth' and line_list[1] == 'sufficient' and line_list[2] == 'pam_unix.so' :               
                    rest = f.read()  
                    f.seek(-len(rest), 1)
                    f.seek(-len(line), 1)
                    f.truncate() 
                    content1 = 'auth        required      pam_faillock.so preauth audit even_deny_root \
deny={}  unlock_time={}\n'.format(user_deny,user_unlocktime)
                    content2 = 'auth        [default=die] pam_faillock.so authfail audit even_deny_root \
deny={}  unlock_time={}\n'.format(user_deny,user_unlocktime)
                    content3 = 'auth        sufficient    pam_faillock.so authsucc audit even_deny_root \
deny={}  unlock_time={}\n'.format(user_deny,user_unlocktime)
                    f.write(content1.encode())
                    f.write(line)
                    f.write(content2.encode())
                    f.write(content3.encode())                                      
                    f.write(rest)
                    break
        user_su_deny = self.ui.checkBox_su.isChecked()
        with open('/etc/pam.d/su', mode='rb+') as f:
            while True:#首先删去已有配置行
                line = f.readline()  # 逐行读取
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                try:
                    if ((line_list[0] == 'auth' and  line_list[1] == 'required'  and line_list[2] == 'pam_wheel.so')
                        or (line_list[0] == '#' and  line_list[1] == 'auth'  and line_list[2] == 'required' and line_list[3] == 'pam_wheel.so' )
                        or (line_list[0] == '#auth' and  line_list[1] == 'required'  and line_list[2] == 'pam_wheel.so' )):                   

                        rest = f.read()  
                        f.seek(-len(rest), 1)
                        f.seek(-len(line), 1) 
                        f.truncate() 
                        f.write(rest)
                        #f.seek(-len(rest), 1)
                except IndexError:
                    continue
            f.seek(0)
            while True:
                #try:
                line = f.readline()  # 逐行读取
                # except IndexError:  # 超出范围则退出
                #     break
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                if line_list[0] == 'auth' and line_list[1] == 'substack' and line_list[2] == 'system-auth' :               
                    rest = f.read()  
                    f.seek(-len(rest), 1)
                    f.truncate() 
                    f.seek(-len(line), 1)
                    if user_su_deny:
                        content = 'auth		required	pam_wheel.so use_uid\n'
                    else:
                        content = '# auth		required	pam_wheel.so use_uid\n'
                    f.write(content.encode())     
                    f.write(line)                                
                    f.write(rest)
                    break


    def sshlogin_conf(self):#如果有这三行的话，就修改；没有的话就在pam_unix上下添加
        print("sshlogin---")
        ssh_deny = self.ui.ssh_deny.value()
        ssh_unlocktime = self.ui.ssh_unlocktime.value()
        with open('/etc/pam.d/password-auth', mode='rb+') as f:
            while True:#首先删去已有配置行
                # try:
                line = f.readline()  # 逐行读取
                # except IndexError:  # 超出范围则退出
                #     break
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                try:
                    if ((line_list[0] == 'auth' and  line_list[1] == 'required'  and line_list[2] == 'pam_faillock.so' and line_list[3] == 'preauth')
                        or (line_list[0] == 'auth' and  line_list[1] == '[default=die]'  and line_list[2] == 'pam_faillock.so' and line_list[3] == 'authfail' )
                        or (line_list[0] == 'auth' and  line_list[1] == 'sufficient'  and line_list[2] == 'pam_faillock.so' and line_list[3] == 'authsucc' )):                   
                        #删除这一句
                        rest = f.read()  
                        f.seek(-len(rest), 1)
                        f.seek(-len(line), 1) 
                        f.truncate() 
                        f.write(rest)
                        f.seek(-len(rest), 1)
                except IndexError:
                    continue
            f.seek(0)
            while True:
                #try:
                line = f.readline()  # 逐行读取
                # except IndexError:  # 超出范围则退出
                #     break
                line_str = line.decode().splitlines()
                if line_str == []:
                    break
                elif line_str == ['']:
                    continue
                line_list = line_str[0].split()
                if line_list[0] == 'auth' and line_list[1] == 'sufficient' and line_list[2] == 'pam_unix.so' :               
                    rest = f.read()  
                    f.seek(-len(rest), 1)
                    f.seek(-len(line), 1)
                    f.truncate() 
                    content1 = 'auth        required      pam_faillock.so preauth audit even_deny_root \
deny={}  unlock_time={}\n'.format(ssh_deny,ssh_unlocktime)
                    content2 = 'auth        [default=die] pam_faillock.so authfail audit even_deny_root \
deny={}  unlock_time={}\n'.format(ssh_deny,ssh_unlocktime)
                    content3 = 'auth        sufficient    pam_faillock.so authsucc audit even_deny_root \
deny={}  unlock_time={}\n'.format(ssh_deny,ssh_unlocktime)
                    f.write(content1.encode())
                    f.write(line)
                    f.write(content2.encode())
                    f.write(content3.encode())                                      
                    f.write(rest)
                    break
                
                    


                
                    

        



if  __name__ == '__main__' :
    myapp = QApplication(sys.argv)
    mywindow = Mainpam()
    mywindow.show()
    sys.exit(myapp.exec_())
