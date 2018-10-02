#import <Foundation/Foundation.h>

extern void ReceiveURL(char*);

@interface GoPasser : NSObject
+ (void)handleGetURLEvent:(NSAppleEventDescriptor *)event;
@end

void StartURLHandler(void);
