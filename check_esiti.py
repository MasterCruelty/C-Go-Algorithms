import requests
import time
url = input("Inserisci il sito di algoritmi: ")
esame = input("Inserisci la prova di cui controllare(ad esempio 20 giugno): ")
print("Inizio a controllare...")
check = False
while check == False:
     resp = requests.get(url)
     if "esiti della prova del "+str(esame) in resp.text:
             print("Esiti pubblicati")
             check = True
             time.sleep(7200)
     else:
             print("Ancora niente esiti")
             time.sleep(60)
